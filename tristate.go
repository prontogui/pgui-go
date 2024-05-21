package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type TristateWith struct {
	Label string
	State int
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

	label field.String
	state field.Integer
}

func (tri *Tristate) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {
	tri.AttachField("Label", &tri.label, pkey, PKeyIndexDontCare, onset)
	tri.AttachField("State", &tri.state, pkey, PKeyIndexDontCare, onset)
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
