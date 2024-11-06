// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A check provides a yes/no, on/off, 1/0, kind of choice to the user.
// It is often represented with a check box like you would see on a form.
type CheckWith struct {
	Checked    bool
	Embodiment string
	Label      string
	Tag        string
}

// Makes a new Check with specified field values.
func (w CheckWith) Make() *Check {
	check := &Check{}
	check.checked.Set(w.Checked)
	check.embodiment.Set(CanonizeEmbodiment(w.Embodiment))
	check.label.Set(w.Label)
	check.tag.Set(w.Tag)
	return check
}

// A check provides a yes/no, on/off, 1/0, kind of choice to the user.
// It is often represented with a check box like you would see on a form.
type Check struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	checked    BooleanField
	embodiment StringField
	label      StringField
	tag        StringField
}

// Creates a new Check and assigns a label.
func NewCheck(label string) *Check {
	return CheckWith{Label: label}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (check *Check) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	check.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Checked, &check.checked},
			{key.FKey_Embodiment, &check.embodiment},
			{key.FKey_Label, &check.label},
			{key.FKey_Tag, &check.tag},
		}
	})
}

// Returns a string representation of this primitive:  the label.
// Implements of fmt:Stringer interface.
func (check *Check) String() string {
	return check.label.Get()
}

// Returns true if the check state is Yes, On, 1, etc., and false if the check state is No, Off, 0, etc.
func (check *Check) Checked() bool {
	return check.checked.Get()
}

// Sets the check state.
func (check *Check) SetChecked(b bool) *Check {
	check.checked.Set(b)
	return check
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (check *Check) Embodiment() string {
	return check.embodiment.Get()
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (check *Check) SetEmbodiment(s string) *Check {
	check.embodiment.Set(CanonizeEmbodiment(s))
	return check
}

// Returns the label to display in the check.
func (check *Check) Label() string {
	return check.label.Get()
}

// Sets the label to display in the check.
func (check *Check) SetLabel(s string) *Check {
	check.label.Set(s)
	return check
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Checks as Table cells.
func (check *Check) Tag() string {
	return check.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Checks as Table cells.
func (check *Check) SetTag(s string) *Check {
	check.tag.Set(s)
	return check
}
