// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/annotation"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/tasklog"
	"github.com/m-mizutani/alertchain/types"
)

// TaskLogCreate is the builder for creating a TaskLog entity.
type TaskLogCreate struct {
	config
	mutation *TaskLogMutation
	hooks    []Hook
}

// SetTaskName sets the "task_name" field.
func (tlc *TaskLogCreate) SetTaskName(s string) *TaskLogCreate {
	tlc.mutation.SetTaskName(s)
	return tlc
}

// SetOptional sets the "optional" field.
func (tlc *TaskLogCreate) SetOptional(b bool) *TaskLogCreate {
	tlc.mutation.SetOptional(b)
	return tlc
}

// SetNillableOptional sets the "optional" field if the given value is not nil.
func (tlc *TaskLogCreate) SetNillableOptional(b *bool) *TaskLogCreate {
	if b != nil {
		tlc.SetOptional(*b)
	}
	return tlc
}

// SetStage sets the "stage" field.
func (tlc *TaskLogCreate) SetStage(i int64) *TaskLogCreate {
	tlc.mutation.SetStage(i)
	return tlc
}

// SetStartedAt sets the "started_at" field.
func (tlc *TaskLogCreate) SetStartedAt(i int64) *TaskLogCreate {
	tlc.mutation.SetStartedAt(i)
	return tlc
}

// SetExitedAt sets the "exited_at" field.
func (tlc *TaskLogCreate) SetExitedAt(i int64) *TaskLogCreate {
	tlc.mutation.SetExitedAt(i)
	return tlc
}

// SetNillableExitedAt sets the "exited_at" field if the given value is not nil.
func (tlc *TaskLogCreate) SetNillableExitedAt(i *int64) *TaskLogCreate {
	if i != nil {
		tlc.SetExitedAt(*i)
	}
	return tlc
}

// SetLog sets the "log" field.
func (tlc *TaskLogCreate) SetLog(s string) *TaskLogCreate {
	tlc.mutation.SetLog(s)
	return tlc
}

// SetNillableLog sets the "log" field if the given value is not nil.
func (tlc *TaskLogCreate) SetNillableLog(s *string) *TaskLogCreate {
	if s != nil {
		tlc.SetLog(*s)
	}
	return tlc
}

// SetErrmsg sets the "errmsg" field.
func (tlc *TaskLogCreate) SetErrmsg(s string) *TaskLogCreate {
	tlc.mutation.SetErrmsg(s)
	return tlc
}

// SetNillableErrmsg sets the "errmsg" field if the given value is not nil.
func (tlc *TaskLogCreate) SetNillableErrmsg(s *string) *TaskLogCreate {
	if s != nil {
		tlc.SetErrmsg(*s)
	}
	return tlc
}

// SetStatus sets the "status" field.
func (tlc *TaskLogCreate) SetStatus(ts types.TaskStatus) *TaskLogCreate {
	tlc.mutation.SetStatus(ts)
	return tlc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tlc *TaskLogCreate) SetNillableStatus(ts *types.TaskStatus) *TaskLogCreate {
	if ts != nil {
		tlc.SetStatus(*ts)
	}
	return tlc
}

// AddAnnotatedIDs adds the "annotated" edge to the Annotation entity by IDs.
func (tlc *TaskLogCreate) AddAnnotatedIDs(ids ...int) *TaskLogCreate {
	tlc.mutation.AddAnnotatedIDs(ids...)
	return tlc
}

// AddAnnotated adds the "annotated" edges to the Annotation entity.
func (tlc *TaskLogCreate) AddAnnotated(a ...*Annotation) *TaskLogCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tlc.AddAnnotatedIDs(ids...)
}

// Mutation returns the TaskLogMutation object of the builder.
func (tlc *TaskLogCreate) Mutation() *TaskLogMutation {
	return tlc.mutation
}

// Save creates the TaskLog in the database.
func (tlc *TaskLogCreate) Save(ctx context.Context) (*TaskLog, error) {
	var (
		err  error
		node *TaskLog
	)
	tlc.defaults()
	if len(tlc.hooks) == 0 {
		if err = tlc.check(); err != nil {
			return nil, err
		}
		node, err = tlc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tlc.check(); err != nil {
				return nil, err
			}
			tlc.mutation = mutation
			if node, err = tlc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tlc.hooks) - 1; i >= 0; i-- {
			if tlc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tlc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tlc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tlc *TaskLogCreate) SaveX(ctx context.Context) *TaskLog {
	v, err := tlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tlc *TaskLogCreate) Exec(ctx context.Context) error {
	_, err := tlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlc *TaskLogCreate) ExecX(ctx context.Context) {
	if err := tlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tlc *TaskLogCreate) defaults() {
	if _, ok := tlc.mutation.Optional(); !ok {
		v := tasklog.DefaultOptional
		tlc.mutation.SetOptional(v)
	}
	if _, ok := tlc.mutation.Status(); !ok {
		v := tasklog.DefaultStatus
		tlc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tlc *TaskLogCreate) check() error {
	if _, ok := tlc.mutation.TaskName(); !ok {
		return &ValidationError{Name: "task_name", err: errors.New(`ent: missing required field "task_name"`)}
	}
	if _, ok := tlc.mutation.Optional(); !ok {
		return &ValidationError{Name: "optional", err: errors.New(`ent: missing required field "optional"`)}
	}
	if _, ok := tlc.mutation.Stage(); !ok {
		return &ValidationError{Name: "stage", err: errors.New(`ent: missing required field "stage"`)}
	}
	if _, ok := tlc.mutation.StartedAt(); !ok {
		return &ValidationError{Name: "started_at", err: errors.New(`ent: missing required field "started_at"`)}
	}
	if _, ok := tlc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	return nil
}

func (tlc *TaskLogCreate) sqlSave(ctx context.Context) (*TaskLog, error) {
	_node, _spec := tlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tlc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tlc *TaskLogCreate) createSpec() (*TaskLog, *sqlgraph.CreateSpec) {
	var (
		_node = &TaskLog{config: tlc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tasklog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tasklog.FieldID,
			},
		}
	)
	if value, ok := tlc.mutation.TaskName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasklog.FieldTaskName,
		})
		_node.TaskName = value
	}
	if value, ok := tlc.mutation.Optional(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: tasklog.FieldOptional,
		})
		_node.Optional = value
	}
	if value, ok := tlc.mutation.Stage(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tasklog.FieldStage,
		})
		_node.Stage = value
	}
	if value, ok := tlc.mutation.StartedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tasklog.FieldStartedAt,
		})
		_node.StartedAt = value
	}
	if value, ok := tlc.mutation.ExitedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tasklog.FieldExitedAt,
		})
		_node.ExitedAt = value
	}
	if value, ok := tlc.mutation.Log(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasklog.FieldLog,
		})
		_node.Log = value
	}
	if value, ok := tlc.mutation.Errmsg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasklog.FieldErrmsg,
		})
		_node.Errmsg = value
	}
	if value, ok := tlc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasklog.FieldStatus,
		})
		_node.Status = value
	}
	if nodes := tlc.mutation.AnnotatedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasklog.AnnotatedTable,
			Columns: []string{tasklog.AnnotatedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: annotation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TaskLogCreateBulk is the builder for creating many TaskLog entities in bulk.
type TaskLogCreateBulk struct {
	config
	builders []*TaskLogCreate
}

// Save creates the TaskLog entities in the database.
func (tlcb *TaskLogCreateBulk) Save(ctx context.Context) ([]*TaskLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tlcb.builders))
	nodes := make([]*TaskLog, len(tlcb.builders))
	mutators := make([]Mutator, len(tlcb.builders))
	for i := range tlcb.builders {
		func(i int, root context.Context) {
			builder := tlcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tlcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tlcb *TaskLogCreateBulk) SaveX(ctx context.Context) []*TaskLog {
	v, err := tlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tlcb *TaskLogCreateBulk) Exec(ctx context.Context) error {
	_, err := tlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlcb *TaskLogCreateBulk) ExecX(ctx context.Context) {
	if err := tlcb.Exec(ctx); err != nil {
		panic(err)
	}
}