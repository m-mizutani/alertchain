// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/m-mizutani/alertchain/pkg/domain/types"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/actionlog"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/job"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/predicate"
)

// ActionLogUpdate is the builder for updating ActionLog entities.
type ActionLogUpdate struct {
	config
	hooks    []Hook
	mutation *ActionLogMutation
}

// Where appends a list predicates to the ActionLogUpdate builder.
func (alu *ActionLogUpdate) Where(ps ...predicate.ActionLog) *ActionLogUpdate {
	alu.mutation.Where(ps...)
	return alu
}

// SetStoppedAt sets the "stopped_at" field.
func (alu *ActionLogUpdate) SetStoppedAt(i int64) *ActionLogUpdate {
	alu.mutation.ResetStoppedAt()
	alu.mutation.SetStoppedAt(i)
	return alu
}

// SetNillableStoppedAt sets the "stopped_at" field if the given value is not nil.
func (alu *ActionLogUpdate) SetNillableStoppedAt(i *int64) *ActionLogUpdate {
	if i != nil {
		alu.SetStoppedAt(*i)
	}
	return alu
}

// AddStoppedAt adds i to the "stopped_at" field.
func (alu *ActionLogUpdate) AddStoppedAt(i int64) *ActionLogUpdate {
	alu.mutation.AddStoppedAt(i)
	return alu
}

// ClearStoppedAt clears the value of the "stopped_at" field.
func (alu *ActionLogUpdate) ClearStoppedAt() *ActionLogUpdate {
	alu.mutation.ClearStoppedAt()
	return alu
}

// SetLog sets the "log" field.
func (alu *ActionLogUpdate) SetLog(s string) *ActionLogUpdate {
	alu.mutation.SetLog(s)
	return alu
}

// SetNillableLog sets the "log" field if the given value is not nil.
func (alu *ActionLogUpdate) SetNillableLog(s *string) *ActionLogUpdate {
	if s != nil {
		alu.SetLog(*s)
	}
	return alu
}

// ClearLog clears the value of the "log" field.
func (alu *ActionLogUpdate) ClearLog() *ActionLogUpdate {
	alu.mutation.ClearLog()
	return alu
}

// SetErrmsg sets the "errmsg" field.
func (alu *ActionLogUpdate) SetErrmsg(s string) *ActionLogUpdate {
	alu.mutation.SetErrmsg(s)
	return alu
}

// SetNillableErrmsg sets the "errmsg" field if the given value is not nil.
func (alu *ActionLogUpdate) SetNillableErrmsg(s *string) *ActionLogUpdate {
	if s != nil {
		alu.SetErrmsg(*s)
	}
	return alu
}

// ClearErrmsg clears the value of the "errmsg" field.
func (alu *ActionLogUpdate) ClearErrmsg() *ActionLogUpdate {
	alu.mutation.ClearErrmsg()
	return alu
}

// SetErrValues sets the "err_values" field.
func (alu *ActionLogUpdate) SetErrValues(s []string) *ActionLogUpdate {
	alu.mutation.SetErrValues(s)
	return alu
}

// ClearErrValues clears the value of the "err_values" field.
func (alu *ActionLogUpdate) ClearErrValues() *ActionLogUpdate {
	alu.mutation.ClearErrValues()
	return alu
}

// SetStackTrace sets the "stack_trace" field.
func (alu *ActionLogUpdate) SetStackTrace(s []string) *ActionLogUpdate {
	alu.mutation.SetStackTrace(s)
	return alu
}

// ClearStackTrace clears the value of the "stack_trace" field.
func (alu *ActionLogUpdate) ClearStackTrace() *ActionLogUpdate {
	alu.mutation.ClearStackTrace()
	return alu
}

// SetStatus sets the "status" field.
func (alu *ActionLogUpdate) SetStatus(ts types.ExecStatus) *ActionLogUpdate {
	alu.mutation.SetStatus(ts)
	return alu
}

// SetJobID sets the "job" edge to the Job entity by ID.
func (alu *ActionLogUpdate) SetJobID(id int) *ActionLogUpdate {
	alu.mutation.SetJobID(id)
	return alu
}

// SetNillableJobID sets the "job" edge to the Job entity by ID if the given value is not nil.
func (alu *ActionLogUpdate) SetNillableJobID(id *int) *ActionLogUpdate {
	if id != nil {
		alu = alu.SetJobID(*id)
	}
	return alu
}

// SetJob sets the "job" edge to the Job entity.
func (alu *ActionLogUpdate) SetJob(j *Job) *ActionLogUpdate {
	return alu.SetJobID(j.ID)
}

// Mutation returns the ActionLogMutation object of the builder.
func (alu *ActionLogUpdate) Mutation() *ActionLogMutation {
	return alu.mutation
}

// ClearJob clears the "job" edge to the Job entity.
func (alu *ActionLogUpdate) ClearJob() *ActionLogUpdate {
	alu.mutation.ClearJob()
	return alu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (alu *ActionLogUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(alu.hooks) == 0 {
		affected, err = alu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActionLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			alu.mutation = mutation
			affected, err = alu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(alu.hooks) - 1; i >= 0; i-- {
			if alu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = alu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, alu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (alu *ActionLogUpdate) SaveX(ctx context.Context) int {
	affected, err := alu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (alu *ActionLogUpdate) Exec(ctx context.Context) error {
	_, err := alu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alu *ActionLogUpdate) ExecX(ctx context.Context) {
	if err := alu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (alu *ActionLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   actionlog.Table,
			Columns: actionlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: actionlog.FieldID,
			},
		},
	}
	if ps := alu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := alu.mutation.StoppedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: actionlog.FieldStoppedAt,
		})
	}
	if value, ok := alu.mutation.AddedStoppedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: actionlog.FieldStoppedAt,
		})
	}
	if alu.mutation.StoppedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: actionlog.FieldStoppedAt,
		})
	}
	if value, ok := alu.mutation.Log(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: actionlog.FieldLog,
		})
	}
	if alu.mutation.LogCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: actionlog.FieldLog,
		})
	}
	if value, ok := alu.mutation.Errmsg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: actionlog.FieldErrmsg,
		})
	}
	if alu.mutation.ErrmsgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: actionlog.FieldErrmsg,
		})
	}
	if value, ok := alu.mutation.ErrValues(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: actionlog.FieldErrValues,
		})
	}
	if alu.mutation.ErrValuesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: actionlog.FieldErrValues,
		})
	}
	if value, ok := alu.mutation.StackTrace(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: actionlog.FieldStackTrace,
		})
	}
	if alu.mutation.StackTraceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: actionlog.FieldStackTrace,
		})
	}
	if value, ok := alu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: actionlog.FieldStatus,
		})
	}
	if alu.mutation.JobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   actionlog.JobTable,
			Columns: []string{actionlog.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := alu.mutation.JobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   actionlog.JobTable,
			Columns: []string{actionlog.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, alu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{actionlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ActionLogUpdateOne is the builder for updating a single ActionLog entity.
type ActionLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ActionLogMutation
}

// SetStoppedAt sets the "stopped_at" field.
func (aluo *ActionLogUpdateOne) SetStoppedAt(i int64) *ActionLogUpdateOne {
	aluo.mutation.ResetStoppedAt()
	aluo.mutation.SetStoppedAt(i)
	return aluo
}

// SetNillableStoppedAt sets the "stopped_at" field if the given value is not nil.
func (aluo *ActionLogUpdateOne) SetNillableStoppedAt(i *int64) *ActionLogUpdateOne {
	if i != nil {
		aluo.SetStoppedAt(*i)
	}
	return aluo
}

// AddStoppedAt adds i to the "stopped_at" field.
func (aluo *ActionLogUpdateOne) AddStoppedAt(i int64) *ActionLogUpdateOne {
	aluo.mutation.AddStoppedAt(i)
	return aluo
}

// ClearStoppedAt clears the value of the "stopped_at" field.
func (aluo *ActionLogUpdateOne) ClearStoppedAt() *ActionLogUpdateOne {
	aluo.mutation.ClearStoppedAt()
	return aluo
}

// SetLog sets the "log" field.
func (aluo *ActionLogUpdateOne) SetLog(s string) *ActionLogUpdateOne {
	aluo.mutation.SetLog(s)
	return aluo
}

// SetNillableLog sets the "log" field if the given value is not nil.
func (aluo *ActionLogUpdateOne) SetNillableLog(s *string) *ActionLogUpdateOne {
	if s != nil {
		aluo.SetLog(*s)
	}
	return aluo
}

// ClearLog clears the value of the "log" field.
func (aluo *ActionLogUpdateOne) ClearLog() *ActionLogUpdateOne {
	aluo.mutation.ClearLog()
	return aluo
}

// SetErrmsg sets the "errmsg" field.
func (aluo *ActionLogUpdateOne) SetErrmsg(s string) *ActionLogUpdateOne {
	aluo.mutation.SetErrmsg(s)
	return aluo
}

// SetNillableErrmsg sets the "errmsg" field if the given value is not nil.
func (aluo *ActionLogUpdateOne) SetNillableErrmsg(s *string) *ActionLogUpdateOne {
	if s != nil {
		aluo.SetErrmsg(*s)
	}
	return aluo
}

// ClearErrmsg clears the value of the "errmsg" field.
func (aluo *ActionLogUpdateOne) ClearErrmsg() *ActionLogUpdateOne {
	aluo.mutation.ClearErrmsg()
	return aluo
}

// SetErrValues sets the "err_values" field.
func (aluo *ActionLogUpdateOne) SetErrValues(s []string) *ActionLogUpdateOne {
	aluo.mutation.SetErrValues(s)
	return aluo
}

// ClearErrValues clears the value of the "err_values" field.
func (aluo *ActionLogUpdateOne) ClearErrValues() *ActionLogUpdateOne {
	aluo.mutation.ClearErrValues()
	return aluo
}

// SetStackTrace sets the "stack_trace" field.
func (aluo *ActionLogUpdateOne) SetStackTrace(s []string) *ActionLogUpdateOne {
	aluo.mutation.SetStackTrace(s)
	return aluo
}

// ClearStackTrace clears the value of the "stack_trace" field.
func (aluo *ActionLogUpdateOne) ClearStackTrace() *ActionLogUpdateOne {
	aluo.mutation.ClearStackTrace()
	return aluo
}

// SetStatus sets the "status" field.
func (aluo *ActionLogUpdateOne) SetStatus(ts types.ExecStatus) *ActionLogUpdateOne {
	aluo.mutation.SetStatus(ts)
	return aluo
}

// SetJobID sets the "job" edge to the Job entity by ID.
func (aluo *ActionLogUpdateOne) SetJobID(id int) *ActionLogUpdateOne {
	aluo.mutation.SetJobID(id)
	return aluo
}

// SetNillableJobID sets the "job" edge to the Job entity by ID if the given value is not nil.
func (aluo *ActionLogUpdateOne) SetNillableJobID(id *int) *ActionLogUpdateOne {
	if id != nil {
		aluo = aluo.SetJobID(*id)
	}
	return aluo
}

// SetJob sets the "job" edge to the Job entity.
func (aluo *ActionLogUpdateOne) SetJob(j *Job) *ActionLogUpdateOne {
	return aluo.SetJobID(j.ID)
}

// Mutation returns the ActionLogMutation object of the builder.
func (aluo *ActionLogUpdateOne) Mutation() *ActionLogMutation {
	return aluo.mutation
}

// ClearJob clears the "job" edge to the Job entity.
func (aluo *ActionLogUpdateOne) ClearJob() *ActionLogUpdateOne {
	aluo.mutation.ClearJob()
	return aluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aluo *ActionLogUpdateOne) Select(field string, fields ...string) *ActionLogUpdateOne {
	aluo.fields = append([]string{field}, fields...)
	return aluo
}

// Save executes the query and returns the updated ActionLog entity.
func (aluo *ActionLogUpdateOne) Save(ctx context.Context) (*ActionLog, error) {
	var (
		err  error
		node *ActionLog
	)
	if len(aluo.hooks) == 0 {
		node, err = aluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActionLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aluo.mutation = mutation
			node, err = aluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aluo.hooks) - 1; i >= 0; i-- {
			if aluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aluo *ActionLogUpdateOne) SaveX(ctx context.Context) *ActionLog {
	node, err := aluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aluo *ActionLogUpdateOne) Exec(ctx context.Context) error {
	_, err := aluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aluo *ActionLogUpdateOne) ExecX(ctx context.Context) {
	if err := aluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aluo *ActionLogUpdateOne) sqlSave(ctx context.Context) (_node *ActionLog, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   actionlog.Table,
			Columns: actionlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: actionlog.FieldID,
			},
		},
	}
	id, ok := aluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ActionLog.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := aluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, actionlog.FieldID)
		for _, f := range fields {
			if !actionlog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != actionlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aluo.mutation.StoppedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: actionlog.FieldStoppedAt,
		})
	}
	if value, ok := aluo.mutation.AddedStoppedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: actionlog.FieldStoppedAt,
		})
	}
	if aluo.mutation.StoppedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: actionlog.FieldStoppedAt,
		})
	}
	if value, ok := aluo.mutation.Log(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: actionlog.FieldLog,
		})
	}
	if aluo.mutation.LogCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: actionlog.FieldLog,
		})
	}
	if value, ok := aluo.mutation.Errmsg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: actionlog.FieldErrmsg,
		})
	}
	if aluo.mutation.ErrmsgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: actionlog.FieldErrmsg,
		})
	}
	if value, ok := aluo.mutation.ErrValues(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: actionlog.FieldErrValues,
		})
	}
	if aluo.mutation.ErrValuesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: actionlog.FieldErrValues,
		})
	}
	if value, ok := aluo.mutation.StackTrace(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: actionlog.FieldStackTrace,
		})
	}
	if aluo.mutation.StackTraceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: actionlog.FieldStackTrace,
		})
	}
	if value, ok := aluo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: actionlog.FieldStatus,
		})
	}
	if aluo.mutation.JobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   actionlog.JobTable,
			Columns: []string{actionlog.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aluo.mutation.JobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   actionlog.JobTable,
			Columns: []string{actionlog.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ActionLog{config: aluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{actionlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
