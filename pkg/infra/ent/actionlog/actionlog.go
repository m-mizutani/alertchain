// Code generated by entc, DO NOT EDIT.

package actionlog

const (
	// Label holds the string label denoting the actionlog type in the database.
	Label = "action_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeArgument holds the string denoting the argument edge name in mutations.
	EdgeArgument = "argument"
	// EdgeExecLogs holds the string denoting the exec_logs edge name in mutations.
	EdgeExecLogs = "exec_logs"
	// Table holds the table name of the actionlog in the database.
	Table = "action_logs"
	// ArgumentTable is the table that holds the argument relation/edge.
	ArgumentTable = "attributes"
	// ArgumentInverseTable is the table name for the Attribute entity.
	// It exists in this package in order to avoid circular dependency with the "attribute" package.
	ArgumentInverseTable = "attributes"
	// ArgumentColumn is the table column denoting the argument relation/edge.
	ArgumentColumn = "action_log_argument"
	// ExecLogsTable is the table that holds the exec_logs relation/edge.
	ExecLogsTable = "exec_logs"
	// ExecLogsInverseTable is the table name for the ExecLog entity.
	// It exists in this package in order to avoid circular dependency with the "execlog" package.
	ExecLogsInverseTable = "exec_logs"
	// ExecLogsColumn is the table column denoting the exec_logs relation/edge.
	ExecLogsColumn = "action_log_exec_logs"
)

// Columns holds all SQL columns for actionlog fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "action_logs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"alert_action_logs",
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