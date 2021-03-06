// Code generated by entc, DO NOT EDIT.

package execlog

import (
	"entgo.io/ent/dialect/sql"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/predicate"
	"github.com/m-mizutani/alertchain/types"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v int64) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// Log applies equality check predicate on the "log" field. It's identical to LogEQ.
func Log(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLog), v))
	})
}

// Errmsg applies equality check predicate on the "errmsg" field. It's identical to ErrmsgEQ.
func Errmsg(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldErrmsg), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v types.ExecStatus) predicate.ExecLog {
	vc := string(v)
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), vc))
	})
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v int64) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v int64) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTimestamp), v))
	})
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...int64) predicate.ExecLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExecLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTimestamp), v...))
	})
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...int64) predicate.ExecLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExecLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTimestamp), v...))
	})
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v int64) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTimestamp), v))
	})
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v int64) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTimestamp), v))
	})
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v int64) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTimestamp), v))
	})
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v int64) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTimestamp), v))
	})
}

// LogEQ applies the EQ predicate on the "log" field.
func LogEQ(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLog), v))
	})
}

// LogNEQ applies the NEQ predicate on the "log" field.
func LogNEQ(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLog), v))
	})
}

// LogIn applies the In predicate on the "log" field.
func LogIn(vs ...string) predicate.ExecLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExecLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLog), v...))
	})
}

// LogNotIn applies the NotIn predicate on the "log" field.
func LogNotIn(vs ...string) predicate.ExecLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExecLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLog), v...))
	})
}

// LogGT applies the GT predicate on the "log" field.
func LogGT(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLog), v))
	})
}

// LogGTE applies the GTE predicate on the "log" field.
func LogGTE(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLog), v))
	})
}

// LogLT applies the LT predicate on the "log" field.
func LogLT(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLog), v))
	})
}

// LogLTE applies the LTE predicate on the "log" field.
func LogLTE(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLog), v))
	})
}

// LogContains applies the Contains predicate on the "log" field.
func LogContains(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLog), v))
	})
}

// LogHasPrefix applies the HasPrefix predicate on the "log" field.
func LogHasPrefix(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLog), v))
	})
}

// LogHasSuffix applies the HasSuffix predicate on the "log" field.
func LogHasSuffix(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLog), v))
	})
}

// LogIsNil applies the IsNil predicate on the "log" field.
func LogIsNil() predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLog)))
	})
}

// LogNotNil applies the NotNil predicate on the "log" field.
func LogNotNil() predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLog)))
	})
}

// LogEqualFold applies the EqualFold predicate on the "log" field.
func LogEqualFold(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLog), v))
	})
}

// LogContainsFold applies the ContainsFold predicate on the "log" field.
func LogContainsFold(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLog), v))
	})
}

// ErrmsgEQ applies the EQ predicate on the "errmsg" field.
func ErrmsgEQ(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldErrmsg), v))
	})
}

// ErrmsgNEQ applies the NEQ predicate on the "errmsg" field.
func ErrmsgNEQ(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldErrmsg), v))
	})
}

// ErrmsgIn applies the In predicate on the "errmsg" field.
func ErrmsgIn(vs ...string) predicate.ExecLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExecLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldErrmsg), v...))
	})
}

// ErrmsgNotIn applies the NotIn predicate on the "errmsg" field.
func ErrmsgNotIn(vs ...string) predicate.ExecLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExecLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldErrmsg), v...))
	})
}

// ErrmsgGT applies the GT predicate on the "errmsg" field.
func ErrmsgGT(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldErrmsg), v))
	})
}

// ErrmsgGTE applies the GTE predicate on the "errmsg" field.
func ErrmsgGTE(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldErrmsg), v))
	})
}

// ErrmsgLT applies the LT predicate on the "errmsg" field.
func ErrmsgLT(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldErrmsg), v))
	})
}

// ErrmsgLTE applies the LTE predicate on the "errmsg" field.
func ErrmsgLTE(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldErrmsg), v))
	})
}

// ErrmsgContains applies the Contains predicate on the "errmsg" field.
func ErrmsgContains(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldErrmsg), v))
	})
}

// ErrmsgHasPrefix applies the HasPrefix predicate on the "errmsg" field.
func ErrmsgHasPrefix(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldErrmsg), v))
	})
}

// ErrmsgHasSuffix applies the HasSuffix predicate on the "errmsg" field.
func ErrmsgHasSuffix(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldErrmsg), v))
	})
}

// ErrmsgIsNil applies the IsNil predicate on the "errmsg" field.
func ErrmsgIsNil() predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldErrmsg)))
	})
}

// ErrmsgNotNil applies the NotNil predicate on the "errmsg" field.
func ErrmsgNotNil() predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldErrmsg)))
	})
}

// ErrmsgEqualFold applies the EqualFold predicate on the "errmsg" field.
func ErrmsgEqualFold(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldErrmsg), v))
	})
}

// ErrmsgContainsFold applies the ContainsFold predicate on the "errmsg" field.
func ErrmsgContainsFold(v string) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldErrmsg), v))
	})
}

// ErrValuesIsNil applies the IsNil predicate on the "err_values" field.
func ErrValuesIsNil() predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldErrValues)))
	})
}

// ErrValuesNotNil applies the NotNil predicate on the "err_values" field.
func ErrValuesNotNil() predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldErrValues)))
	})
}

// StackTraceIsNil applies the IsNil predicate on the "stack_trace" field.
func StackTraceIsNil() predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStackTrace)))
	})
}

// StackTraceNotNil applies the NotNil predicate on the "stack_trace" field.
func StackTraceNotNil() predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStackTrace)))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v types.ExecStatus) predicate.ExecLog {
	vc := string(v)
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), vc))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v types.ExecStatus) predicate.ExecLog {
	vc := string(v)
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), vc))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...types.ExecStatus) predicate.ExecLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.ExecLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...types.ExecStatus) predicate.ExecLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.ExecLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v types.ExecStatus) predicate.ExecLog {
	vc := string(v)
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), vc))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v types.ExecStatus) predicate.ExecLog {
	vc := string(v)
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), vc))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v types.ExecStatus) predicate.ExecLog {
	vc := string(v)
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), vc))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v types.ExecStatus) predicate.ExecLog {
	vc := string(v)
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), vc))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v types.ExecStatus) predicate.ExecLog {
	vc := string(v)
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), vc))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v types.ExecStatus) predicate.ExecLog {
	vc := string(v)
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), vc))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v types.ExecStatus) predicate.ExecLog {
	vc := string(v)
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), vc))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v types.ExecStatus) predicate.ExecLog {
	vc := string(v)
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), vc))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v types.ExecStatus) predicate.ExecLog {
	vc := string(v)
	return predicate.ExecLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), vc))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ExecLog) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ExecLog) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
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
func Not(p predicate.ExecLog) predicate.ExecLog {
	return predicate.ExecLog(func(s *sql.Selector) {
		p(s.Not())
	})
}
