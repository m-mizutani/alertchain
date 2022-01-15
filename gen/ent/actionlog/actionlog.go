// Code generated by entc, DO NOT EDIT.

package actionlog

const (
	// Label holds the string label denoting the actionlog type in the database.
	Label = "action_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldStoppedAt holds the string denoting the stopped_at field in the database.
	FieldStoppedAt = "stopped_at"
	// FieldLog holds the string denoting the log field in the database.
	FieldLog = "log"
	// FieldErrmsg holds the string denoting the errmsg field in the database.
	FieldErrmsg = "errmsg"
	// FieldErrValues holds the string denoting the err_values field in the database.
	FieldErrValues = "err_values"
	// FieldStackTrace holds the string denoting the stack_trace field in the database.
	FieldStackTrace = "stack_trace"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeJob holds the string denoting the job edge name in mutations.
	EdgeJob = "job"
	// Table holds the table name of the actionlog in the database.
	Table = "action_logs"
	// JobTable is the table that holds the job relation/edge.
	JobTable = "action_logs"
	// JobInverseTable is the table name for the Job entity.
	// It exists in this package in order to avoid circular dependency with the "job" package.
	JobInverseTable = "jobs"
	// JobColumn is the table column denoting the job relation/edge.
	JobColumn = "job_action_logs"
)

// Columns holds all SQL columns for actionlog fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldStartedAt,
	FieldStoppedAt,
	FieldLog,
	FieldErrmsg,
	FieldErrValues,
	FieldStackTrace,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "action_logs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"job_action_logs",
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