// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/alert"
	"github.com/m-mizutani/alertchain/types"
)

// Alert is the model entity for the Alert schema.
type Alert struct {
	config `json:"-"`
	// ID of the ent.
	ID types.AlertID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Detector holds the value of the "detector" field.
	Detector string `json:"detector,omitempty"`
	// Status holds the value of the "status" field.
	Status types.AlertStatus `json:"status,omitempty"`
	// Severity holds the value of the "severity" field.
	Severity types.Severity `json:"severity,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ClosedAt holds the value of the "closed_at" field.
	ClosedAt *time.Time `json:"closed_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AlertQuery when eager-loading is set.
	Edges AlertEdges `json:"edges"`
}

// AlertEdges holds the relations/edges for other nodes in the graph.
type AlertEdges struct {
	// Attributes holds the value of the attributes edge.
	Attributes []*Attribute `json:"attributes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AttributesOrErr returns the Attributes value or an error if the edge
// was not loaded in eager-loading.
func (e AlertEdges) AttributesOrErr() ([]*Attribute, error) {
	if e.loadedTypes[0] {
		return e.Attributes, nil
	}
	return nil, &NotLoadedError{edge: "attributes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Alert) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case alert.FieldID, alert.FieldTitle, alert.FieldDescription, alert.FieldDetector, alert.FieldStatus, alert.FieldSeverity:
			values[i] = new(sql.NullString)
		case alert.FieldCreatedAt, alert.FieldClosedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Alert", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Alert fields.
func (a *Alert) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case alert.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				a.ID = types.AlertID(value.String)
			}
		case alert.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		case alert.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		case alert.FieldDetector:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detector", values[i])
			} else if value.Valid {
				a.Detector = value.String
			}
		case alert.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				a.Status = types.AlertStatus(value.String)
			}
		case alert.FieldSeverity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field severity", values[i])
			} else if value.Valid {
				a.Severity = types.Severity(value.String)
			}
		case alert.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case alert.FieldClosedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field closed_at", values[i])
			} else if value.Valid {
				a.ClosedAt = new(time.Time)
				*a.ClosedAt = value.Time
			}
		}
	}
	return nil
}

// QueryAttributes queries the "attributes" edge of the Alert entity.
func (a *Alert) QueryAttributes() *AttributeQuery {
	return (&AlertClient{config: a.config}).QueryAttributes(a)
}

// Update returns a builder for updating this Alert.
// Note that you need to call Alert.Unwrap() before calling this method if this Alert
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Alert) Update() *AlertUpdateOne {
	return (&AlertClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Alert entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Alert) Unwrap() *Alert {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Alert is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Alert) String() string {
	var builder strings.Builder
	builder.WriteString("Alert(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", title=")
	builder.WriteString(a.Title)
	builder.WriteString(", description=")
	builder.WriteString(a.Description)
	builder.WriteString(", detector=")
	builder.WriteString(a.Detector)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", a.Status))
	builder.WriteString(", severity=")
	builder.WriteString(fmt.Sprintf("%v", a.Severity))
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	if v := a.ClosedAt; v != nil {
		builder.WriteString(", closed_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Alerts is a parsable slice of Alert.
type Alerts []*Alert

func (a Alerts) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
