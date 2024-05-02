package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type CheckWith struct {
	Label   string
	Checked bool
}

// Makes a new Check with specified field values.
func (w CheckWith) Make() *Check {
	cmd := &Check{}
	cmd.label.Set(w.Label)
	cmd.checked.Set(w.Checked)
	return cmd
}

type Check struct {
	// Mix-in the common guts for primitives (B-side fields, implements primitive interface, etc.)
	Reserved

	label   field.String
	checked field.Boolean
	changed field.Event
}

func (check *Check) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	check.AttachField("Label", &check.label)
	check.AttachField("Checked", &check.checked)
	check.AttachField("Changed", &check.changed)

	// Prepare all fields for updates
	check.Reserved.PrepareForUpdates(pkey, onset)
}

func (check *Check) Label() string {
	return check.label.Get()
}

func (check *Check) SetLabel(s string) {
	check.label.Set(s)
}

func (check *Check) Checked() bool {
	return check.checked.Get()
}

func (check *Check) SetChecked(b bool) {
	check.checked.Set(b)
}

func (check *Check) Changed() bool {
	return check.changed.Get()
}
