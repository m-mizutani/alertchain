// Code generated by entc, DO NOT EDIT.

package attribute

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/m-mizutani/alertchain/pkg/domain/types"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKey), v))
	})
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v types.AttrType) predicate.Attribute {
	vc := string(v)
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), vc))
	})
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKey), v))
	})
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKey), v))
	})
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.Attribute {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attribute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldKey), v...))
	})
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.Attribute {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attribute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldKey), v...))
	})
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKey), v))
	})
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKey), v))
	})
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKey), v))
	})
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKey), v))
	})
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldKey), v))
	})
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldKey), v))
	})
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldKey), v))
	})
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldKey), v))
	})
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldKey), v))
	})
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValue), v))
	})
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.Attribute {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attribute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldValue), v...))
	})
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.Attribute {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Attribute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldValue), v...))
	})
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValue), v))
	})
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValue), v))
	})
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValue), v))
	})
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValue), v))
	})
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldValue), v))
	})
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldValue), v))
	})
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldValue), v))
	})
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldValue), v))
	})
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldValue), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v types.AttrType) predicate.Attribute {
	vc := string(v)
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), vc))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v types.AttrType) predicate.Attribute {
	vc := string(v)
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), vc))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...types.AttrType) predicate.Attribute {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Attribute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...types.AttrType) predicate.Attribute {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Attribute(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v types.AttrType) predicate.Attribute {
	vc := string(v)
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), vc))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v types.AttrType) predicate.Attribute {
	vc := string(v)
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), vc))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v types.AttrType) predicate.Attribute {
	vc := string(v)
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), vc))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v types.AttrType) predicate.Attribute {
	vc := string(v)
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), vc))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v types.AttrType) predicate.Attribute {
	vc := string(v)
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), vc))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v types.AttrType) predicate.Attribute {
	vc := string(v)
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), vc))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v types.AttrType) predicate.Attribute {
	vc := string(v)
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), vc))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v types.AttrType) predicate.Attribute {
	vc := string(v)
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), vc))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v types.AttrType) predicate.Attribute {
	vc := string(v)
	return predicate.Attribute(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), vc))
	})
}

// HasAnnotations applies the HasEdge predicate on the "annotations" edge.
func HasAnnotations() predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AnnotationsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AnnotationsTable, AnnotationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAnnotationsWith applies the HasEdge predicate on the "annotations" edge with a given conditions (other predicates).
func HasAnnotationsWith(preds ...predicate.Annotation) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AnnotationsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AnnotationsTable, AnnotationsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAlert applies the HasEdge predicate on the "alert" edge.
func HasAlert() predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AlertTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AlertTable, AlertColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAlertWith applies the HasEdge predicate on the "alert" edge with a given conditions (other predicates).
func HasAlertWith(preds ...predicate.Alert) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AlertInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AlertTable, AlertColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Attribute) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Attribute) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Attribute) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		p(s.Not())
	})
}
