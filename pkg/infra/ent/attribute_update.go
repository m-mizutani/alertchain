// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/alert"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/annotation"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/attribute"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/predicate"
	"github.com/m-mizutani/alertchain/types"
)

// AttributeUpdate is the builder for updating Attribute entities.
type AttributeUpdate struct {
	config
	hooks    []Hook
	mutation *AttributeMutation
}

// Where appends a list predicates to the AttributeUpdate builder.
func (au *AttributeUpdate) Where(ps ...predicate.Attribute) *AttributeUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetKey sets the "key" field.
func (au *AttributeUpdate) SetKey(s string) *AttributeUpdate {
	au.mutation.SetKey(s)
	return au
}

// SetValue sets the "value" field.
func (au *AttributeUpdate) SetValue(s string) *AttributeUpdate {
	au.mutation.SetValue(s)
	return au
}

// SetType sets the "type" field.
func (au *AttributeUpdate) SetType(tt types.AttrType) *AttributeUpdate {
	au.mutation.SetType(tt)
	return au
}

// SetContext sets the "context" field.
func (au *AttributeUpdate) SetContext(s []string) *AttributeUpdate {
	au.mutation.SetContext(s)
	return au
}

// AddAnnotationIDs adds the "annotations" edge to the Annotation entity by IDs.
func (au *AttributeUpdate) AddAnnotationIDs(ids ...int) *AttributeUpdate {
	au.mutation.AddAnnotationIDs(ids...)
	return au
}

// AddAnnotations adds the "annotations" edges to the Annotation entity.
func (au *AttributeUpdate) AddAnnotations(a ...*Annotation) *AttributeUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddAnnotationIDs(ids...)
}

// SetAlertID sets the "alert" edge to the Alert entity by ID.
func (au *AttributeUpdate) SetAlertID(id types.AlertID) *AttributeUpdate {
	au.mutation.SetAlertID(id)
	return au
}

// SetNillableAlertID sets the "alert" edge to the Alert entity by ID if the given value is not nil.
func (au *AttributeUpdate) SetNillableAlertID(id *types.AlertID) *AttributeUpdate {
	if id != nil {
		au = au.SetAlertID(*id)
	}
	return au
}

// SetAlert sets the "alert" edge to the Alert entity.
func (au *AttributeUpdate) SetAlert(a *Alert) *AttributeUpdate {
	return au.SetAlertID(a.ID)
}

// Mutation returns the AttributeMutation object of the builder.
func (au *AttributeUpdate) Mutation() *AttributeMutation {
	return au.mutation
}

// ClearAnnotations clears all "annotations" edges to the Annotation entity.
func (au *AttributeUpdate) ClearAnnotations() *AttributeUpdate {
	au.mutation.ClearAnnotations()
	return au
}

// RemoveAnnotationIDs removes the "annotations" edge to Annotation entities by IDs.
func (au *AttributeUpdate) RemoveAnnotationIDs(ids ...int) *AttributeUpdate {
	au.mutation.RemoveAnnotationIDs(ids...)
	return au
}

// RemoveAnnotations removes "annotations" edges to Annotation entities.
func (au *AttributeUpdate) RemoveAnnotations(a ...*Annotation) *AttributeUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveAnnotationIDs(ids...)
}

// ClearAlert clears the "alert" edge to the Alert entity.
func (au *AttributeUpdate) ClearAlert() *AttributeUpdate {
	au.mutation.ClearAlert()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AttributeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttributeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AttributeUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AttributeUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AttributeUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AttributeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   attribute.Table,
			Columns: attribute.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: attribute.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: attribute.FieldKey,
		})
	}
	if value, ok := au.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: attribute.FieldValue,
		})
	}
	if value, ok := au.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: attribute.FieldType,
		})
	}
	if value, ok := au.mutation.Context(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: attribute.FieldContext,
		})
	}
	if au.mutation.AnnotationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AnnotationsTable,
			Columns: []string{attribute.AnnotationsColumn},
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
	if nodes := au.mutation.RemovedAnnotationsIDs(); len(nodes) > 0 && !au.mutation.AnnotationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AnnotationsTable,
			Columns: []string{attribute.AnnotationsColumn},
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
	if nodes := au.mutation.AnnotationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AnnotationsTable,
			Columns: []string{attribute.AnnotationsColumn},
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
	if au.mutation.AlertCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   attribute.AlertTable,
			Columns: []string{attribute.AlertColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: alert.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AlertIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   attribute.AlertTable,
			Columns: []string{attribute.AlertColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: alert.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attribute.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AttributeUpdateOne is the builder for updating a single Attribute entity.
type AttributeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AttributeMutation
}

// SetKey sets the "key" field.
func (auo *AttributeUpdateOne) SetKey(s string) *AttributeUpdateOne {
	auo.mutation.SetKey(s)
	return auo
}

// SetValue sets the "value" field.
func (auo *AttributeUpdateOne) SetValue(s string) *AttributeUpdateOne {
	auo.mutation.SetValue(s)
	return auo
}

// SetType sets the "type" field.
func (auo *AttributeUpdateOne) SetType(tt types.AttrType) *AttributeUpdateOne {
	auo.mutation.SetType(tt)
	return auo
}

// SetContext sets the "context" field.
func (auo *AttributeUpdateOne) SetContext(s []string) *AttributeUpdateOne {
	auo.mutation.SetContext(s)
	return auo
}

// AddAnnotationIDs adds the "annotations" edge to the Annotation entity by IDs.
func (auo *AttributeUpdateOne) AddAnnotationIDs(ids ...int) *AttributeUpdateOne {
	auo.mutation.AddAnnotationIDs(ids...)
	return auo
}

// AddAnnotations adds the "annotations" edges to the Annotation entity.
func (auo *AttributeUpdateOne) AddAnnotations(a ...*Annotation) *AttributeUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddAnnotationIDs(ids...)
}

// SetAlertID sets the "alert" edge to the Alert entity by ID.
func (auo *AttributeUpdateOne) SetAlertID(id types.AlertID) *AttributeUpdateOne {
	auo.mutation.SetAlertID(id)
	return auo
}

// SetNillableAlertID sets the "alert" edge to the Alert entity by ID if the given value is not nil.
func (auo *AttributeUpdateOne) SetNillableAlertID(id *types.AlertID) *AttributeUpdateOne {
	if id != nil {
		auo = auo.SetAlertID(*id)
	}
	return auo
}

// SetAlert sets the "alert" edge to the Alert entity.
func (auo *AttributeUpdateOne) SetAlert(a *Alert) *AttributeUpdateOne {
	return auo.SetAlertID(a.ID)
}

// Mutation returns the AttributeMutation object of the builder.
func (auo *AttributeUpdateOne) Mutation() *AttributeMutation {
	return auo.mutation
}

// ClearAnnotations clears all "annotations" edges to the Annotation entity.
func (auo *AttributeUpdateOne) ClearAnnotations() *AttributeUpdateOne {
	auo.mutation.ClearAnnotations()
	return auo
}

// RemoveAnnotationIDs removes the "annotations" edge to Annotation entities by IDs.
func (auo *AttributeUpdateOne) RemoveAnnotationIDs(ids ...int) *AttributeUpdateOne {
	auo.mutation.RemoveAnnotationIDs(ids...)
	return auo
}

// RemoveAnnotations removes "annotations" edges to Annotation entities.
func (auo *AttributeUpdateOne) RemoveAnnotations(a ...*Annotation) *AttributeUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveAnnotationIDs(ids...)
}

// ClearAlert clears the "alert" edge to the Alert entity.
func (auo *AttributeUpdateOne) ClearAlert() *AttributeUpdateOne {
	auo.mutation.ClearAlert()
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AttributeUpdateOne) Select(field string, fields ...string) *AttributeUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Attribute entity.
func (auo *AttributeUpdateOne) Save(ctx context.Context) (*Attribute, error) {
	var (
		err  error
		node *Attribute
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttributeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AttributeUpdateOne) SaveX(ctx context.Context) *Attribute {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AttributeUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AttributeUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AttributeUpdateOne) sqlSave(ctx context.Context) (_node *Attribute, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   attribute.Table,
			Columns: attribute.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: attribute.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Attribute.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attribute.FieldID)
		for _, f := range fields {
			if !attribute.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != attribute.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: attribute.FieldKey,
		})
	}
	if value, ok := auo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: attribute.FieldValue,
		})
	}
	if value, ok := auo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: attribute.FieldType,
		})
	}
	if value, ok := auo.mutation.Context(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: attribute.FieldContext,
		})
	}
	if auo.mutation.AnnotationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AnnotationsTable,
			Columns: []string{attribute.AnnotationsColumn},
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
	if nodes := auo.mutation.RemovedAnnotationsIDs(); len(nodes) > 0 && !auo.mutation.AnnotationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AnnotationsTable,
			Columns: []string{attribute.AnnotationsColumn},
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
	if nodes := auo.mutation.AnnotationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AnnotationsTable,
			Columns: []string{attribute.AnnotationsColumn},
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
	if auo.mutation.AlertCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   attribute.AlertTable,
			Columns: []string{attribute.AlertColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: alert.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AlertIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   attribute.AlertTable,
			Columns: []string{attribute.AlertColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: alert.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Attribute{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attribute.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
