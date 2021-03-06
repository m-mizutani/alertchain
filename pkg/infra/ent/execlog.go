// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/execlog"
	"github.com/m-mizutani/alertchain/types"
)

// ExecLog is the model entity for the ExecLog schema.
type ExecLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Log holds the value of the "log" field.
	Log string `json:"log,omitempty"`
	// Errmsg holds the value of the "errmsg" field.
	Errmsg string `json:"errmsg,omitempty"`
	// ErrValues holds the value of the "err_values" field.
	ErrValues []string `json:"err_values,omitempty"`
	// StackTrace holds the value of the "stack_trace" field.
	StackTrace []string `json:"stack_trace,omitempty"`
	// Status holds the value of the "status" field.
	Status               types.ExecStatus `json:"status,omitempty"`
	action_log_exec_logs *int
	task_log_exec_logs   *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ExecLog) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case execlog.FieldErrValues, execlog.FieldStackTrace:
			values[i] = new([]byte)
		case execlog.FieldID, execlog.FieldTimestamp:
			values[i] = new(sql.NullInt64)
		case execlog.FieldLog, execlog.FieldErrmsg, execlog.FieldStatus:
			values[i] = new(sql.NullString)
		case execlog.ForeignKeys[0]: // action_log_exec_logs
			values[i] = new(sql.NullInt64)
		case execlog.ForeignKeys[1]: // task_log_exec_logs
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ExecLog", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ExecLog fields.
func (el *ExecLog) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case execlog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			el.ID = int(value.Int64)
		case execlog.FieldTimestamp:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				el.Timestamp = value.Int64
			}
		case execlog.FieldLog:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field log", values[i])
			} else if value.Valid {
				el.Log = value.String
			}
		case execlog.FieldErrmsg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field errmsg", values[i])
			} else if value.Valid {
				el.Errmsg = value.String
			}
		case execlog.FieldErrValues:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field err_values", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &el.ErrValues); err != nil {
					return fmt.Errorf("unmarshal field err_values: %w", err)
				}
			}
		case execlog.FieldStackTrace:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field stack_trace", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &el.StackTrace); err != nil {
					return fmt.Errorf("unmarshal field stack_trace: %w", err)
				}
			}
		case execlog.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				el.Status = types.ExecStatus(value.String)
			}
		case execlog.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field action_log_exec_logs", value)
			} else if value.Valid {
				el.action_log_exec_logs = new(int)
				*el.action_log_exec_logs = int(value.Int64)
			}
		case execlog.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field task_log_exec_logs", value)
			} else if value.Valid {
				el.task_log_exec_logs = new(int)
				*el.task_log_exec_logs = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ExecLog.
// Note that you need to call ExecLog.Unwrap() before calling this method if this ExecLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (el *ExecLog) Update() *ExecLogUpdateOne {
	return (&ExecLogClient{config: el.config}).UpdateOne(el)
}

// Unwrap unwraps the ExecLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (el *ExecLog) Unwrap() *ExecLog {
	tx, ok := el.config.driver.(*txDriver)
	if !ok {
		panic("ent: ExecLog is not a transactional entity")
	}
	el.config.driver = tx.drv
	return el
}

// String implements the fmt.Stringer.
func (el *ExecLog) String() string {
	var builder strings.Builder
	builder.WriteString("ExecLog(")
	builder.WriteString(fmt.Sprintf("id=%v", el.ID))
	builder.WriteString(", timestamp=")
	builder.WriteString(fmt.Sprintf("%v", el.Timestamp))
	builder.WriteString(", log=")
	builder.WriteString(el.Log)
	builder.WriteString(", errmsg=")
	builder.WriteString(el.Errmsg)
	builder.WriteString(", err_values=")
	builder.WriteString(fmt.Sprintf("%v", el.ErrValues))
	builder.WriteString(", stack_trace=")
	builder.WriteString(fmt.Sprintf("%v", el.StackTrace))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", el.Status))
	builder.WriteByte(')')
	return builder.String()
}

// ExecLogs is a parsable slice of ExecLog.
type ExecLogs []*ExecLog

func (el ExecLogs) config(cfg config) {
	for _i := range el {
		el[_i].config = cfg
	}
}
