package usecase_test

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/m-mizutani/alertchain"
	"github.com/m-mizutani/alertchain/pkg/infra"
	"github.com/m-mizutani/alertchain/pkg/infra/db"
	"github.com/m-mizutani/alertchain/pkg/infra/ent"
	"github.com/m-mizutani/alertchain/pkg/usecase"
	"github.com/m-mizutani/alertchain/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type sleeper struct {
	Done bool
}

func (x *sleeper) Name() string        { return "sleeper" }
func (x *sleeper) Description() string { return "sleep random duration" }
func (x *sleeper) Execute(ctx context.Context, alert *alertchain.Alert) error {
	time.Sleep(time.Millisecond * time.Duration(rand.Int31n(2000)))
	x.Done = true
	return nil
}

func (x *sleeper) Optionable(alert *alertchain.Alert) bool {
	panic("not implemented") // TODO: Implement
}

func setupAlertTest(t *testing.T) (usecase.Interface, infra.Clients, *alertchain.Chain) {
	chain := &alertchain.Chain{}

	clients := infra.Clients{
		DB: db.NewDBMock(t),
	}
	uc := usecase.New(clients, chain)

	return uc, clients, chain
}

type mock struct {
	Exec func(alert *alertchain.Alert) error
}

func (x *mock) Name() string                            { return "mock" }
func (x *mock) Description() string                     { return "mocking" }
func (x *mock) Optionable(alert *alertchain.Alert) bool { return false }
func (x *mock) Execute(ctx context.Context, alert *alertchain.Alert) error {
	return x.Exec(alert)
}

func TestRecvAlert(t *testing.T) {
	uc, clients, chain := setupAlertTest(t)

	var done bool
	chain.NewStage().AddTask(&mock{
		Exec: func(alert *alertchain.Alert) error {
			alert.UpdateSeverity(types.SevAffected)
			alert.UpdateStatus(types.StatusClosed)
			if err := alert.Commit(context.Background()); err != nil {
				return err
			}
			done = true
			return nil
		},
	})

	input := alertchain.Alert{
		Alert: ent.Alert{
			Title:    "five",
			Detector: "blue",
		},
	}
	ctx, wg := usecase.ContextWithWaitGroup(context.Background())
	alert, err := uc.RecvAlert(ctx, &input)
	require.NoError(t, err)
	require.NotNil(t, alert)

	wg.Wait()
	assert.True(t, done)

	got, err := clients.DB.GetAlert(context.Background(), alert.ID)
	require.NoError(t, err)
	assert.Equal(t, alert.Title, got.Title)
	assert.Equal(t, types.SevAffected, got.Severity)
	assert.Equal(t, types.StatusClosed, got.Status)
}

func TestRecvAlertDoNotUpdate(t *testing.T) {
	t.Run("do not update severity and status by overwriting vars", func(t *testing.T) {
		uc, clients, chain := setupAlertTest(t)

		var done bool
		chain.NewStage().AddTask(&mock{
			Exec: func(alert *alertchain.Alert) error {
				alert.Severity = types.SevAffected
				alert.Status = types.StatusClosed
				if err := alert.Commit(context.Background()); err != nil {
					return err
				}
				done = true
				return nil
			},
		})

		input := alertchain.Alert{
			Alert: ent.Alert{
				Title:    "five",
				Detector: "blue",
			},
		}
		ctx, wg := usecase.ContextWithWaitGroup(context.Background())
		alert, err := uc.RecvAlert(ctx, &input)
		require.NoError(t, err)
		require.NotNil(t, alert)

		wg.Wait()
		assert.True(t, done)

		got, err := clients.DB.GetAlert(context.Background(), alert.ID)
		require.NoError(t, err)
		assert.Equal(t, alert.Title, got.Title)
		assert.NotEqual(t, types.SevAffected, got.Severity)
		assert.NotEqual(t, types.StatusClosed, got.Status)
	})
}

func TestRecvAlertMultipleTask(t *testing.T) {
	uc, _, chain := setupAlertTest(t)

	stage := chain.NewStage()
	stage.Timeout = time.Second
	for i := 0; i < 32; i++ {
		stage.AddTask(&sleeper{})
	}

	ctx, wg := usecase.ContextWithWaitGroup(context.Background())
	input := alertchain.Alert{
		Alert: ent.Alert{
			Title:    "five",
			Detector: "blue",
		},
	}
	_, err := uc.RecvAlert(ctx, &input)
	require.NoError(t, err)
	wg.Wait()

	require.Len(t, stage.Tasks, 32)
	for _, task := range stage.Tasks {
		s, ok := task.(*sleeper)
		require.True(t, ok)
		assert.True(t, s.Done)
	}
}

func TestRecvAlertMassiveAnnotation(t *testing.T) {
	const multiplex = 12

	// Do not use sqlite3 because of table lock error
	passwd := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	if passwd == "" || dbName == "" {
		t.Skip("MYSQL_ROOT_PASSWORD and MYSQL_DATABASE are required")
	}

	chain := &alertchain.Chain{}
	dsn := fmt.Sprintf("root:%s@tcp(localhost:3306)/%s", passwd, dbName)
	client, err := db.New("mysql", dsn)
	require.NoError(t, err)

	clients := infra.Clients{DB: client}
	uc := usecase.New(clients, chain)

	stage := chain.NewStage()
	stage.Timeout = time.Second
	for i := 0; i < multiplex; i++ {
		stage.AddTask(&mock{
			Exec: func(alert *alertchain.Alert) error {
				require.Len(t, alert.Attributes, 1)
				alert.Attributes[0].Annotate(&alertchain.Annotation{
					Annotation: ent.Annotation{
						Source:    "x",
						Timestamp: rand.Int63(), /* nosec */
						Name:      "y",
						Value:     "z",
					},
				})
				return nil
			},
		})
	}

	ctx, wg := usecase.ContextWithWaitGroup(context.Background())
	input := alertchain.Alert{
		Alert: ent.Alert{
			Title:    "five",
			Detector: "blue",
		},
		Attributes: []*alertchain.Attribute{
			{
				Attribute: ent.Attribute{
					Key:   "color",
					Value: "red",
					Type:  types.AttrUserID,
				},
			},
		},
	}
	created, err := uc.RecvAlert(ctx, &input)
	require.NoError(t, err)
	wg.Wait()

	alert, err := clients.DB.GetAlert(context.Background(), created.Alert.ID)
	require.NoError(t, err)
	require.Len(t, alert.Edges.Attributes[0].Edges.Annotations, multiplex)
	for _, ann := range alert.Edges.Attributes[0].Edges.Annotations {
		assert.Equal(t, "x", ann.Source)
		assert.Equal(t, "y", ann.Name)
		assert.Equal(t, "z", ann.Value)
		assert.Greater(t, ann.Timestamp, int64(0))
	}
}

func TestRecvAlertErrorHandling(t *testing.T) {
	t.Run("exit on error", func(t *testing.T) {
		uc, _, chain := setupAlertTest(t)

		stage := chain.NewStage()
		stage.ExitOnErr = true
		stage.AddTask(&mock{
			Exec: func(alert *alertchain.Alert) error { return nil },
		})
		stage.AddTask(&mock{
			Exec: func(alert *alertchain.Alert) error { return errors.New("bomb!") },
		})

		done2ndStage := false
		chain.NewStage().AddTask(&mock{
			Exec: func(alert *alertchain.Alert) error {
				done2ndStage = true
				return nil
			},
		})

		input := alertchain.Alert{
			Alert: ent.Alert{
				Title:    "five",
				Detector: "blue",
			},
		}
		ctx, wg := usecase.ContextWithWaitGroup(context.Background())
		_, err := uc.RecvAlert(ctx, &input)
		require.NoError(t, err)
		wg.Wait()
		assert.False(t, done2ndStage)
	})

	t.Run("not exit on error", func(t *testing.T) {
		uc, _, chain := setupAlertTest(t)

		stage := chain.NewStage()
		// Default: stage.ExitOnErr = false
		stage.AddTask(&mock{
			Exec: func(alert *alertchain.Alert) error { return nil },
		})
		stage.AddTask(&mock{
			Exec: func(alert *alertchain.Alert) error { return errors.New("bomb!") },
		})

		done2ndStage := false
		chain.NewStage().AddTask(&mock{
			Exec: func(alert *alertchain.Alert) error {
				done2ndStage = true
				return nil
			},
		})

		input := alertchain.Alert{
			Alert: ent.Alert{
				Title:    "five",
				Detector: "blue",
			},
		}
		ctx, wg := usecase.ContextWithWaitGroup(context.Background())
		_, err := uc.RecvAlert(ctx, &input)
		require.NoError(t, err)
		wg.Wait()
		assert.True(t, done2ndStage)
	})
}
