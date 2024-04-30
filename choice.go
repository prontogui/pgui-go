package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type Choice struct {
	// Mix-in the common guts for primitives (B-side fields, implements primitive interface, etc.)
	Reserved

	choice  field.String
	issued  field.Boolean
	choices field.Strings1D
}

func (choice *Choice) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	choice.AttachField("Choice", &choice.choice)
	choice.AttachField("Issued", &choice.issued)
	choice.AttachField("Choices", &choice.choices)

	// Prepare all fields for updates
	choice.Reserved.PrepareForUpdates(pkey, onset)
}

func (choice *Choice) Choice() string {
	return choice.choice.Get()
}

func (choice *Choice) SetChoice(s string) {
	choice.choice.Set(s)
}

func (choice *Choice) Issued() bool {
	return choice.issued.Get()
}

func (choice *Choice) SetIssued(b bool) {
	choice.issued.Set(b)
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
