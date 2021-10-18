// Code generated by entc, DO NOT EDIT.

package alert

import (
	"github.com/m-mizutani/alertchain/types"
)

const (
	// Label holds the string label denoting the alert type in the database.
	Label = "alert"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldDetector holds the string denoting the detector field in the database.
	FieldDetector = "detector"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldSeverity holds the string denoting the severity field in the database.
	FieldSeverity = "severity"
	// FieldDetectedAt holds the string denoting the detected_at field in the database.
	FieldDetectedAt = "detected_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldClosedAt holds the string denoting the closed_at field in the database.
	FieldClosedAt = "closed_at"
	// EdgeAttributes holds the string denoting the attributes edge name in mutations.
	EdgeAttributes = "attributes"
	// EdgeReferences holds the string denoting the references edge name in mutations.
	EdgeReferences = "references"
	// EdgeTaskLogs holds the string denoting the task_logs edge name in mutations.
	EdgeTaskLogs = "task_logs"
	// EdgeActionLogs holds the string denoting the action_logs edge name in mutations.
	EdgeActionLogs = "action_logs"
	// Table holds the table name of the alert in the database.
	Table = "alerts"
	// AttributesTable is the table that holds the attributes relation/edge.
	AttributesTable = "attributes"
	// AttributesInverseTable is the table name for the Attribute entity.
	// It exists in this package in order to avoid circular dependency with the "attribute" package.
	AttributesInverseTable = "attributes"
	// AttributesColumn is the table column denoting the attributes relation/edge.
	AttributesColumn = "alert_attributes"
	// ReferencesTable is the table that holds the references relation/edge.
	ReferencesTable = "references"
	// ReferencesInverseTable is the table name for the Reference entity.
	// It exists in this package in order to avoid circular dependency with the "reference" package.
	ReferencesInverseTable = "references"
	// ReferencesColumn is the table column denoting the references relation/edge.
	ReferencesColumn = "alert_references"
	// TaskLogsTable is the table that holds the task_logs relation/edge.
	TaskLogsTable = "task_logs"
	// TaskLogsInverseTable is the table name for the TaskLog entity.
	// It exists in this package in order to avoid circular dependency with the "tasklog" package.
	TaskLogsInverseTable = "task_logs"
	// TaskLogsColumn is the table column denoting the task_logs relation/edge.
	TaskLogsColumn = "alert_task_logs"
	// ActionLogsTable is the table that holds the action_logs relation/edge.
	ActionLogsTable = "action_logs"
	// ActionLogsInverseTable is the table name for the ActionLog entity.
	// It exists in this package in order to avoid circular dependency with the "actionlog" package.
	ActionLogsInverseTable = "action_logs"
	// ActionLogsColumn is the table column denoting the action_logs relation/edge.
	ActionLogsColumn = "alert_action_logs"
)

// Columns holds all SQL columns for alert fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
	FieldDetector,
	FieldStatus,
	FieldSeverity,
	FieldDetectedAt,
	FieldCreatedAt,
	FieldClosedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus types.AlertStatus
)
