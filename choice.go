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
	Tag        string
}

// Makes a new Choice with specified field values.
func (w ChoiceWith) Make() *Choice {
	choice := &Choice{}
	choice.choice.Set(w.Choice)
	choice.choices.Set(w.Choices)
	choice.embodiment.Set(CanonizeEmbodiment(w.Embodiment))
	choice.tag.Set(w.Tag)
	return choice
}

// A choice is a user selection from a set of choices.  It is often represented using a pull-down list.
type Choice struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	choice     StringField
	choices    Strings1DField
	embodiment StringField
	tag        StringField
}

// Creates a new Choice and assigns the initiali Choice and Choices fields.
func NewChoice(choices ...string) *Choice {
	return ChoiceWith{Choices: choices}.Make().SetChoiceIndex(0)
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
			{key.FKey_Tag, &choice.tag},
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
	choice.embodiment.Set(CanonizeEmbodiment(s))
	return choice
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Choices as Table cells.
func (choice *Choice) Tag() string {
	return choice.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Choices as Table cells.
func (choice *Choice) SetTag(s string) *Choice {
	choice.tag.Set(s)
	return choice
}

// Returns the index (0, 1, ..) of selected choice or -1 if choice is empty.  This is a covenvenience
// function as an alternative to Choice().  The canonical storage of choice remains a string.
func (choice *Choice) ChoiceIndex() int {
	currentChoice := choice.choice.Get()

	for index, choice := range choice.choices.Get() {
		if choice == currentChoice {
			return index
		}
	}
	return -1
}

// Sets the selected choice or empty if none chosen or if index is out of range.
// This is a covenvenience function as an alternative to SetChoice().  The canonical storage of
// choice remains a string.
func (choice *Choice) SetChoiceIndex(index int) *Choice {

	currentChoices := choice.choices.Get()

	// Valid index?
	if index >= 0 && index < len(currentChoices) {
		choice.choice.Set(currentChoices[index])
	} else {
		choice.choice.Set("")
	}

	return choice
}
