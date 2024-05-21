package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type TristateWith struct {
	Embodiment string
	Label      string
	State      int
}

func (w TristateWith) Make() *Tristate {
	tri := &Tristate{}
	tri.label.Set(w.Label)
	tri.state.Set(w.State)
	return tri
}

type Tristate struct {
	// Mix-in the common guts for primitives
	Reserved

	embodiment field.String
	label      field.String
	state      field.Integer
}

func (tri *Tristate) GetFieldRefs() []FieldRef {
	return []FieldRef{
		{key.FKey_Embodiment, &tri.embodiment},
		{key.FKey_Label, &tri.label},
		{key.FKey_State, &tri.state},
	}
}

func (tri *Tristate) Embodiment() string {
	return tri.embodiment.Get()
}

func (tri *Tristate) SetEmbodiment(s string) {
	tri.embodiment.Set(s)
}

func (tri *Tristate) Label() string {
	return tri.label.Get()
}

func (tri *Tristate) SetLabel(s string) {
	tri.label.Set(s)
}

func (tri *Tristate) State() int {
	return tri.state.Get()
}

func (tri *Tristate) SetState(i int) {
	tri.state.Set(i)
}
