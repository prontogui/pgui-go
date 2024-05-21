package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type CheckWith struct {
	Checked    bool
	Embodiment string
	Label      string
}

// Makes a new Check with specified field values.
func (w CheckWith) Make() *Check {
	check := &Check{}
	check.checked.Set(w.Checked)
	check.embodiment.Set(w.Embodiment)
	check.label.Set(w.Label)
	return check
}

type Check struct {
	// Mix-in the common guts for primitives
	Reserved

	checked    field.Boolean
	embodiment field.String
	label      field.String
}

func (check *Check) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	check.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Checked, &check.checked},
			{key.FKey_Embodiment, &check.embodiment},
			{key.FKey_Label, &check.label},
		}
	})
}

func (check *Check) Checked() bool {
	return check.checked.Get()
}

func (check *Check) SetChecked(b bool) {
	check.checked.Set(b)
}

func (check *Check) Embodiment() string {
	return check.embodiment.Get()
}

func (check *Check) SetEmbodiment(s string) {
	check.embodiment.Set(s)
}

func (check *Check) Label() string {
	return check.label.Get()
}

func (check *Check) SetLabel(s string) {
	check.label.Set(s)
}
