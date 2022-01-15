// Code generated by entc, DO NOT EDIT.

package job

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/m-mizutani/alertchain/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Step applies equality check predicate on the "step" field. It's identical to StepEQ.
func Step(v int64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStep), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// StepEQ applies the EQ predicate on the "step" field.
func StepEQ(v int64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStep), v))
	})
}

// StepNEQ applies the NEQ predicate on the "step" field.
func StepNEQ(v int64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStep), v))
	})
}

// StepIn applies the In predicate on the "step" field.
func StepIn(vs ...int64) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStep), v...))
	})
}

// StepNotIn applies the NotIn predicate on the "step" field.
func StepNotIn(vs ...int64) predicate.Job {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Job(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStep), v...))
	})
}

// StepGT applies the GT predicate on the "step" field.
func StepGT(v int64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStep), v))
	})
}

// StepGTE applies the GTE predicate on the "step" field.
func StepGTE(v int64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStep), v))
	})
}

// StepLT applies the LT predicate on the "step" field.
func StepLT(v int64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStep), v))
	})
}

// StepLTE applies the LTE predicate on the "step" field.
func StepLTE(v int64) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStep), v))
	})
}

// HasAlert applies the HasEdge predicate on the "alert" edge.
func HasAlert() predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AlertTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AlertTable, AlertColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAlertWith applies the HasEdge predicate on the "alert" edge with a given conditions (other predicates).
func HasAlertWith(preds ...predicate.Alert) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
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

// HasActionLogs applies the HasEdge predicate on the "action_logs" edge.
func HasActionLogs() predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActionLogsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ActionLogsTable, ActionLogsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActionLogsWith applies the HasEdge predicate on the "action_logs" edge with a given conditions (other predicates).
func HasActionLogsWith(preds ...predicate.ActionLog) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActionLogsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ActionLogsTable, ActionLogsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Job) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Job) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
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
func Not(p predicate.Job) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		p(s.Not())
	})
}