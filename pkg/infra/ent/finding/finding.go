// Code generated by entc, DO NOT EDIT.

package finding

const (
	// Label holds the string label denoting the finding type in the database.
	Label = "finding"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// Table holds the table name of the finding in the database.
	Table = "findings"
)

// Columns holds all SQL columns for finding fields.
var Columns = []string{
	FieldID,
	FieldTime,
	FieldSource,
	FieldKey,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "findings"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"attribute_findings",
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
