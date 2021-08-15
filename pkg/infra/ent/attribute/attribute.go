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
	// EdgeFindings holds the string denoting the findings edge name in mutations.
	EdgeFindings = "findings"
	// Table holds the table name of the attribute in the database.
	Table = "attributes"
	// FindingsTable is the table that holds the findings relation/edge.
	FindingsTable = "findings"
	// FindingsInverseTable is the table name for the Finding entity.
	// It exists in this package in order to avoid circular dependency with the "finding" package.
	FindingsInverseTable = "findings"
	// FindingsColumn is the table column denoting the findings relation/edge.
	FindingsColumn = "attribute_findings"
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
	"alert_attributes",
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