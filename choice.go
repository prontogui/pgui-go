package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type ChoiceWith struct {
	Choice     string
	Choices    []string
	Embodiment string
}

func (w ChoiceWith) Make() *Choice {
	choice := &Choice{}
	choice.choice.Set(w.Choice)
	choice.choices.Set(w.Choices)
	choice.embodiment.Set(w.Embodiment)
	return choice
}

type Choice struct {
	// Mix-in the common guts for primitives
	Reserved

	choice     field.String
	choices    field.Strings1D
	embodiment field.String
}

func (choice *Choice) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	choice.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Choice, &choice.choice},
			{key.FKey_Choices, &choice.choices},
			{key.FKey_Embodiment, &choice.embodiment},
		}
	})
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

func (choice *Choice) Embodiment() string {
	return choice.embodiment.Get()
}

func (choice *Choice) SetEmbodiment(s string) {
	choice.embodiment.Set(s)
}
