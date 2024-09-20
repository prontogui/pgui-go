// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A choice is a user selection from a set of choices.  It is often represented using a pull-down list.
type ChoiceWith struct {
	Choice     string
	Choices    []string
	Embodiment string
}

// Makes a new Choice with specified field values.
func (w ChoiceWith) Make() *Choice {
	choice := &Choice{}
	choice.choice.Set(w.Choice)
	choice.choices.Set(w.Choices)
	choice.embodiment.Set(w.Embodiment)
	return choice
}

// A choice is a user selection from a set of choices.  It is often represented using a pull-down list.
type Choice struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	choice     StringField
	choices    Strings1DField
	embodiment StringField
}

// Creates a new Choice and assigns the initiali Choice and Choices fields.
func NewChoice(choice string, choices []string) *Choice {
	return ChoiceWith{Choice: choice, Choices: choices}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (choice *Choice) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	choice.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Choice, &choice.choice},
			{key.FKey_Choices, &choice.choices},
			{key.FKey_Embodiment, &choice.embodiment},
		}
	})
}

// Returns a string representation of this primitive:  the current choice.
// Implements of fmt:Stringer interface.
func (choice *Choice) String() string {
	return choice.choice.Get()
}

// Returns the selected choice or empty if none chosen.
func (choice *Choice) Choice() string {
	return choice.choice.Get()
}

// Sets the selected choice or empty if none chosen.
func (choice *Choice) SetChoice(s string) *Choice {
	choice.choice.Set(s)
	return choice
}

// Returns the set of valid choices to choose from.
func (choice *Choice) Choices() []string {
	return choice.choices.Get()
}

// Sets the set of valid choices to choose from.
func (choice *Choice) SetChoices(sa []string) *Choice {
	choice.choices.Set(sa)
	return choice
}

// Set the Choices field using variadic string arguments.
func (choice *Choice) SetChoicesVA(sa ...string) *Choice {
	choice.choices.Set(sa)
	return choice
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (choice *Choice) Embodiment() string {
	return choice.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (choice *Choice) SetEmbodiment(s string) *Choice {
	choice.embodiment.Set(s)
	return choice
}
