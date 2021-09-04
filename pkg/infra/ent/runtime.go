// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/m-mizutani/alertchain/pkg/infra/ent/alert"
	"github.com/m-mizutani/alertchain/pkg/infra/schema"
	"github.com/m-mizutani/alertchain/types"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	alertFields := schema.Alert{}.Fields()
	_ = alertFields
	// alertDescStatus is the schema descriptor for status field.
	alertDescStatus := alertFields[4].Descriptor()
	// alert.DefaultStatus holds the default value on creation for the status field.
	alert.DefaultStatus = types.AlertStatus(alertDescStatus.Default.(string))
}
