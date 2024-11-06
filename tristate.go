// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A choice presented to the user that give three possible states:  Affirmative (Yes, On, 1, etc.),
// Negative (No, Off, 0, etc.), and Indeterminate.
type TristateWith struct {
	Embodiment string
	Label      string
	State      int
	Tag        string
}

// Creates a new TriState using the supplied field assignments.
func (w TristateWith) Make() *Tristate {
	tri := &Tristate{}
	tri.embodiment.Set(CanonizeEmbodiment(w.Embodiment))
	tri.label.Set(w.Label)
	tri.state.Set(w.State)
	tri.tag.Set(w.Tag)
	return tri
}

// A choice presented to the user that give three possible states:  Affirmative (Yes, On, 1, etc.),
// Negative (No, Off, 0, etc.), and Indeterminate.
type Tristate struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	label      StringField
	state      IntegerField
	tag        StringField
}

// Create a new TriState and assign a label.
func NewTristate(label string) *Tristate {
	return TristateWith{Label: label}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (tri *Tristate) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	tri.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &tri.embodiment},
			{key.FKey_Label, &tri.label},
			{key.FKey_State, &tri.state},
			{key.FKey_Tag, &tri.tag},
		}
	})
}

// Returns a string representation of this primitive:  the label.
// Implements of fmt:Stringer interface.
func (tri *Tristate) String() string {
	return tri.label.Get()
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (tri *Tristate) Embodiment() string {
	return tri.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (tri *Tristate) SetEmbodiment(s string) *Tristate {
	tri.embodiment.Set(CanonizeEmbodiment(s))
	return tri
}

// Returns the label to display along with the tristate option.
func (tri *Tristate) Label() string {
	return tri.label.Get()
}

// Sets the label to display along with the tristate option.
func (tri *Tristate) SetLabel(s string) *Tristate {
	tri.label.Set(s)
	return tri
}

// Returns the state of the option (0 = Negative, 1 = Affirmative, and -1 = Indeterminate).
func (tri *Tristate) State() int {
	return tri.state.Get()
}

// Sets the state of the option (0 = Negative, 1 = Affirmative, and -1 = Indeterminate).
func (tri *Tristate) SetState(i int) *Tristate {
	tri.state.Set(i)
	return tri
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Tristates as Table cells.
func (tri *Tristate) Tag() string {
	return tri.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Tristates as Table cells.
func (tri *Tristate) SetTag(s string) *Tristate {
	tri.tag.Set(s)
	return tri
}
