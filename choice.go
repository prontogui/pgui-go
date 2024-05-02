package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type ChoiceWith struct {
	choice  string
	choices []string
}

func (w ChoiceWith) Make() *Choice {
	choice := &Choice{}
	choice.choice.Set(w.choice)
	choice.choices.Set(w.choices)
	return choice
}

type Choice struct {
	// Mix-in the common guts for primitives (B-side fields, implements primitive interface, etc.)
	Reserved

	choice  field.String
	choices field.Strings1D
	changed field.Event
}

func (choice *Choice) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	choice.AttachField("Choice", &choice.choice)
	choice.AttachField("Choices", &choice.choices)
	choice.AttachField("Changed", &choice.changed)

	// Prepare all fields for updates
	choice.Reserved.PrepareForUpdates(pkey, onset)
}

func (choice *Choice) Choice() string {
	return choice.choice.Get()
}

func (choice *Choice) SetChoice(s string) {
	choice.choice.Set(s)
}

func (choice *Choice) Choices() []string {
	return choice.choices.Get()
}

func (choice *Choice) SetChoices(sa []string) {
	choice.choices.Set(sa)
}

// Set the Choices field using variadic string arguments.
func (choice *Choice) SetChoicesVA(sa ...string) {
	choice.choices.Set(sa)
}

func (choice *Choice) Changed() bool {
	return choice.changed.Get()
}
