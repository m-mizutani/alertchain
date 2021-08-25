// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/annotation"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/predicate"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/tasklog"
	"github.com/m-mizutani/alertchain/types"
)

// TaskLogUpdate is the builder for updating TaskLog entities.
type TaskLogUpdate struct {
	config
	hooks    []Hook
	mutation *TaskLogMutation
}

// Where appends a list predicates to the TaskLogUpdate builder.
func (tlu *TaskLogUpdate) Where(ps ...predicate.TaskLog) *TaskLogUpdate {
	tlu.mutation.Where(ps...)
	return tlu
}

// SetExitedAt sets the "exited_at" field.
func (tlu *TaskLogUpdate) SetExitedAt(i int64) *TaskLogUpdate {
	tlu.mutation.ResetExitedAt()
	tlu.mutation.SetExitedAt(i)
	return tlu
}

// SetNillableExitedAt sets the "exited_at" field if the given value is not nil.
func (tlu *TaskLogUpdate) SetNillableExitedAt(i *int64) *TaskLogUpdate {
	if i != nil {
		tlu.SetExitedAt(*i)
	}
	return tlu
}

// AddExitedAt adds i to the "exited_at" field.
func (tlu *TaskLogUpdate) AddExitedAt(i int64) *TaskLogUpdate {
	tlu.mutation.AddExitedAt(i)
	return tlu
}

// ClearExitedAt clears the value of the "exited_at" field.
func (tlu *TaskLogUpdate) ClearExitedAt() *TaskLogUpdate {
	tlu.mutation.ClearExitedAt()
	return tlu
}

// SetLog sets the "log" field.
func (tlu *TaskLogUpdate) SetLog(s string) *TaskLogUpdate {
	tlu.mutation.SetLog(s)
	return tlu
}

// SetNillableLog sets the "log" field if the given value is not nil.
func (tlu *TaskLogUpdate) SetNillableLog(s *string) *TaskLogUpdate {
	if s != nil {
		tlu.SetLog(*s)
	}
	return tlu
}

// ClearLog clears the value of the "log" field.
func (tlu *TaskLogUpdate) ClearLog() *TaskLogUpdate {
	tlu.mutation.ClearLog()
	return tlu
}

// SetErrmsg sets the "errmsg" field.
func (tlu *TaskLogUpdate) SetErrmsg(s string) *TaskLogUpdate {
	tlu.mutation.SetErrmsg(s)
	return tlu
}

// SetNillableErrmsg sets the "errmsg" field if the given value is not nil.
func (tlu *TaskLogUpdate) SetNillableErrmsg(s *string) *TaskLogUpdate {
	if s != nil {
		tlu.SetErrmsg(*s)
	}
	return tlu
}

// ClearErrmsg clears the value of the "errmsg" field.
func (tlu *TaskLogUpdate) ClearErrmsg() *TaskLogUpdate {
	tlu.mutation.ClearErrmsg()
	return tlu
}

// SetStatus sets the "status" field.
func (tlu *TaskLogUpdate) SetStatus(ts types.TaskStatus) *TaskLogUpdate {
	tlu.mutation.SetStatus(ts)
	return tlu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tlu *TaskLogUpdate) SetNillableStatus(ts *types.TaskStatus) *TaskLogUpdate {
	if ts != nil {
		tlu.SetStatus(*ts)
	}
	return tlu
}

// AddAnnotatedIDs adds the "annotated" edge to the Annotation entity by IDs.
func (tlu *TaskLogUpdate) AddAnnotatedIDs(ids ...int) *TaskLogUpdate {
	tlu.mutation.AddAnnotatedIDs(ids...)
	return tlu
}

// AddAnnotated adds the "annotated" edges to the Annotation entity.
func (tlu *TaskLogUpdate) AddAnnotated(a ...*Annotation) *TaskLogUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tlu.AddAnnotatedIDs(ids...)
}

// Mutation returns the TaskLogMutation object of the builder.
func (tlu *TaskLogUpdate) Mutation() *TaskLogMutation {
	return tlu.mutation
}

// ClearAnnotated clears all "annotated" edges to the Annotation entity.
func (tlu *TaskLogUpdate) ClearAnnotated() *TaskLogUpdate {
	tlu.mutation.ClearAnnotated()
	return tlu
}

// RemoveAnnotatedIDs removes the "annotated" edge to Annotation entities by IDs.
func (tlu *TaskLogUpdate) RemoveAnnotatedIDs(ids ...int) *TaskLogUpdate {
	tlu.mutation.RemoveAnnotatedIDs(ids...)
	return tlu
}

// RemoveAnnotated removes "annotated" edges to Annotation entities.
func (tlu *TaskLogUpdate) RemoveAnnotated(a ...*Annotation) *TaskLogUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tlu.RemoveAnnotatedIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tlu *TaskLogUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tlu.hooks) == 0 {
		affected, err = tlu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tlu.mutation = mutation
			affected, err = tlu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tlu.hooks) - 1; i >= 0; i-- {
			if tlu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tlu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tlu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tlu *TaskLogUpdate) SaveX(ctx context.Context) int {
	affected, err := tlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tlu *TaskLogUpdate) Exec(ctx context.Context) error {
	_, err := tlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlu *TaskLogUpdate) ExecX(ctx context.Context) {
	if err := tlu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tlu *TaskLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tasklog.Table,
			Columns: tasklog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tasklog.FieldID,
			},
		},
	}
	if ps := tlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tlu.mutation.ExitedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tasklog.FieldExitedAt,
		})
	}
	if value, ok := tlu.mutation.AddedExitedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tasklog.FieldExitedAt,
		})
	}
	if tlu.mutation.ExitedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: tasklog.FieldExitedAt,
		})
	}
	if value, ok := tlu.mutation.Log(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasklog.FieldLog,
		})
	}
	if tlu.mutation.LogCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tasklog.FieldLog,
		})
	}
	if value, ok := tlu.mutation.Errmsg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasklog.FieldErrmsg,
		})
	}
	if tlu.mutation.ErrmsgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tasklog.FieldErrmsg,
		})
	}
	if value, ok := tlu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasklog.FieldStatus,
		})
	}
	if tlu.mutation.AnnotatedCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tlu.mutation.RemovedAnnotatedIDs(); len(nodes) > 0 && !tlu.mutation.AnnotatedCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tlu.mutation.AnnotatedIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tasklog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TaskLogUpdateOne is the builder for updating a single TaskLog entity.
type TaskLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskLogMutation
}

// SetExitedAt sets the "exited_at" field.
func (tluo *TaskLogUpdateOne) SetExitedAt(i int64) *TaskLogUpdateOne {
	tluo.mutation.ResetExitedAt()
	tluo.mutation.SetExitedAt(i)
	return tluo
}

// SetNillableExitedAt sets the "exited_at" field if the given value is not nil.
func (tluo *TaskLogUpdateOne) SetNillableExitedAt(i *int64) *TaskLogUpdateOne {
	if i != nil {
		tluo.SetExitedAt(*i)
	}
	return tluo
}

// AddExitedAt adds i to the "exited_at" field.
func (tluo *TaskLogUpdateOne) AddExitedAt(i int64) *TaskLogUpdateOne {
	tluo.mutation.AddExitedAt(i)
	return tluo
}

// ClearExitedAt clears the value of the "exited_at" field.
func (tluo *TaskLogUpdateOne) ClearExitedAt() *TaskLogUpdateOne {
	tluo.mutation.ClearExitedAt()
	return tluo
}

// SetLog sets the "log" field.
func (tluo *TaskLogUpdateOne) SetLog(s string) *TaskLogUpdateOne {
	tluo.mutation.SetLog(s)
	return tluo
}

// SetNillableLog sets the "log" field if the given value is not nil.
func (tluo *TaskLogUpdateOne) SetNillableLog(s *string) *TaskLogUpdateOne {
	if s != nil {
		tluo.SetLog(*s)
	}
	return tluo
}

// ClearLog clears the value of the "log" field.
func (tluo *TaskLogUpdateOne) ClearLog() *TaskLogUpdateOne {
	tluo.mutation.ClearLog()
	return tluo
}

// SetErrmsg sets the "errmsg" field.
func (tluo *TaskLogUpdateOne) SetErrmsg(s string) *TaskLogUpdateOne {
	tluo.mutation.SetErrmsg(s)
	return tluo
}

// SetNillableErrmsg sets the "errmsg" field if the given value is not nil.
func (tluo *TaskLogUpdateOne) SetNillableErrmsg(s *string) *TaskLogUpdateOne {
	if s != nil {
		tluo.SetErrmsg(*s)
	}
	return tluo
}

// ClearErrmsg clears the value of the "errmsg" field.
func (tluo *TaskLogUpdateOne) ClearErrmsg() *TaskLogUpdateOne {
	tluo.mutation.ClearErrmsg()
	return tluo
}

// SetStatus sets the "status" field.
func (tluo *TaskLogUpdateOne) SetStatus(ts types.TaskStatus) *TaskLogUpdateOne {
	tluo.mutation.SetStatus(ts)
	return tluo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tluo *TaskLogUpdateOne) SetNillableStatus(ts *types.TaskStatus) *TaskLogUpdateOne {
	if ts != nil {
		tluo.SetStatus(*ts)
	}
	return tluo
}

// AddAnnotatedIDs adds the "annotated" edge to the Annotation entity by IDs.
func (tluo *TaskLogUpdateOne) AddAnnotatedIDs(ids ...int) *TaskLogUpdateOne {
	tluo.mutation.AddAnnotatedIDs(ids...)
	return tluo
}

// AddAnnotated adds the "annotated" edges to the Annotation entity.
func (tluo *TaskLogUpdateOne) AddAnnotated(a ...*Annotation) *TaskLogUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tluo.AddAnnotatedIDs(ids...)
}

// Mutation returns the TaskLogMutation object of the builder.
func (tluo *TaskLogUpdateOne) Mutation() *TaskLogMutation {
	return tluo.mutation
}

// ClearAnnotated clears all "annotated" edges to the Annotation entity.
func (tluo *TaskLogUpdateOne) ClearAnnotated() *TaskLogUpdateOne {
	tluo.mutation.ClearAnnotated()
	return tluo
}

// RemoveAnnotatedIDs removes the "annotated" edge to Annotation entities by IDs.
func (tluo *TaskLogUpdateOne) RemoveAnnotatedIDs(ids ...int) *TaskLogUpdateOne {
	tluo.mutation.RemoveAnnotatedIDs(ids...)
	return tluo
}

// RemoveAnnotated removes "annotated" edges to Annotation entities.
func (tluo *TaskLogUpdateOne) RemoveAnnotated(a ...*Annotation) *TaskLogUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tluo.RemoveAnnotatedIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tluo *TaskLogUpdateOne) Select(field string, fields ...string) *TaskLogUpdateOne {
	tluo.fields = append([]string{field}, fields...)
	return tluo
}

// Save executes the query and returns the updated TaskLog entity.
func (tluo *TaskLogUpdateOne) Save(ctx context.Context) (*TaskLog, error) {
	var (
		err  error
		node *TaskLog
	)
	if len(tluo.hooks) == 0 {
		node, err = tluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tluo.mutation = mutation
			node, err = tluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tluo.hooks) - 1; i >= 0; i-- {
			if tluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tluo *TaskLogUpdateOne) SaveX(ctx context.Context) *TaskLog {
	node, err := tluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tluo *TaskLogUpdateOne) Exec(ctx context.Context) error {
	_, err := tluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tluo *TaskLogUpdateOne) ExecX(ctx context.Context) {
	if err := tluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tluo *TaskLogUpdateOne) sqlSave(ctx context.Context) (_node *TaskLog, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tasklog.Table,
			Columns: tasklog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tasklog.FieldID,
			},
		},
	}
	id, ok := tluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TaskLog.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tasklog.FieldID)
		for _, f := range fields {
			if !tasklog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tasklog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tluo.mutation.ExitedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tasklog.FieldExitedAt,
		})
	}
	if value, ok := tluo.mutation.AddedExitedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tasklog.FieldExitedAt,
		})
	}
	if tluo.mutation.ExitedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: tasklog.FieldExitedAt,
		})
	}
	if value, ok := tluo.mutation.Log(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasklog.FieldLog,
		})
	}
	if tluo.mutation.LogCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tasklog.FieldLog,
		})
	}
	if value, ok := tluo.mutation.Errmsg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasklog.FieldErrmsg,
		})
	}
	if tluo.mutation.ErrmsgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tasklog.FieldErrmsg,
		})
	}
	if value, ok := tluo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasklog.FieldStatus,
		})
	}
	if tluo.mutation.AnnotatedCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tluo.mutation.RemovedAnnotatedIDs(); len(nodes) > 0 && !tluo.mutation.AnnotatedCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tluo.mutation.AnnotatedIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TaskLog{config: tluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tasklog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}