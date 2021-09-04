// Code generated by entc, DO NOT EDIT.

package attribute

const (
	// Label holds the string label denoting the attribute type in the database.
	Label = "attribute"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldContext holds the string denoting the context field in the database.
	FieldContext = "context"
	// EdgeAnnotations holds the string denoting the annotations edge name in mutations.
	EdgeAnnotations = "annotations"
	// EdgeAlert holds the string denoting the alert edge name in mutations.
	EdgeAlert = "alert"
	// Table holds the table name of the attribute in the database.
	Table = "attributes"
	// AnnotationsTable is the table that holds the annotations relation/edge.
	AnnotationsTable = "annotations"
	// AnnotationsInverseTable is the table name for the Annotation entity.
	// It exists in this package in order to avoid circular dependency with the "annotation" package.
	AnnotationsInverseTable = "annotations"
	// AnnotationsColumn is the table column denoting the annotations relation/edge.
	AnnotationsColumn = "attribute_annotations"
	// AlertTable is the table that holds the alert relation/edge.
	AlertTable = "attributes"
	// AlertInverseTable is the table name for the Alert entity.
	// It exists in this package in order to avoid circular dependency with the "alert" package.
	AlertInverseTable = "alerts"
	// AlertColumn is the table column denoting the alert relation/edge.
	AlertColumn = "attribute_alert"
)

// Columns holds all SQL columns for attribute fields.
var Columns = []string{
	FieldID,
	FieldKey,
	FieldValue,
	FieldType,
	FieldContext,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "attributes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"action_log_argument",
	"alert_attributes",
	"attribute_alert",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
