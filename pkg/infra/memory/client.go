package memory

import (
	"sync"
	"time"

	"github.com/m-mizutani/alertchain/pkg/domain/interfaces"
	"github.com/m-mizutani/alertchain/pkg/domain/model"
	"github.com/m-mizutani/alertchain/pkg/domain/types"
)

type lock struct {
	mutex     sync.Mutex
	expiresAt time.Time
}

type Client struct {
	attrs map[types.Namespace]map[types.AttrID]*model.Attribute
	locks map[types.Namespace]*lock

	attrMutex sync.RWMutex
	lockMutex sync.Mutex
}

func New() *Client {
	return &Client{
		attrs: map[types.Namespace]map[types.AttrID]*model.Attribute{},
		locks: map[types.Namespace]*lock{},
	}
}

// Close implements interfaces.Database.
func (x *Client) Close() error {
	return nil
}

// GetAttrs implements interfaces.Database.
func (x *Client) GetAttrs(ctx *model.Context, ns types.Namespace) (model.Attributes, error) {
	x.attrMutex.RLock()
	defer x.attrMutex.RUnlock()

	attrs, ok := x.attrs[ns]
	if !ok {
		return nil, nil
	}

	var ret model.Attributes
	for _, a := range attrs {
		ret = append(ret, *a)
	}

	return ret, nil
}

// PutAttrs implements interfaces.Database.
func (x *Client) PutAttrs(ctx *model.Context, ns types.Namespace, attrs model.Attributes) error {
	x.attrMutex.Lock()
	if _, ok := x.attrs[ns]; !ok {
		x.attrs[ns] = map[types.AttrID]*model.Attribute{}
	}
	x.attrMutex.Unlock()

	for i, src := range attrs {
		if dst, ok := x.attrs[ns][src.ID]; ok {
			dst.Value = src.Value
			dst.ExpiresAt = src.ExpiresAt
		} else {
			x.attrs[ns][src.ID] = &attrs[i]
		}
	}

	return nil
}

// Lock implements interfaces.Database.
func (x *Client) Lock(ctx *model.Context, ns types.Namespace, timeout time.Time) error {
	locked := make(chan struct{}, 1)
	go func() {
		for i := 0; ; i++ {
			x.lockMutex.Lock()
			if l, ok := x.locks[ns]; !ok || l.expiresAt.Before(time.Now()) {
				x.locks[ns] = &lock{}
			}
			x.lockMutex.Unlock()

			if x.locks[ns].mutex.TryLock() {
				x.locks[ns].expiresAt = timeout
				break
			}

			time.Sleep(10 * time.Millisecond)
		}
		locked <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		// cancelled
	case <-locked:
		// locked
	}

	return nil
}

// Unlock implements interfaces.Database.
func (x *Client) Unlock(ctx *model.Context, ns types.Namespace) error {
	if _, ok := x.locks[ns]; !ok {
		return nil
	}

	x.locks[ns].mutex.Unlock()
	return nil
}

var _ interfaces.Database = (*Client)(nil)
