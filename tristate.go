// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

type TristateWith struct {
	Embodiment string
	Label      string
	State      int
}

func (w TristateWith) Make() *Tristate {
	tri := &Tristate{}
	tri.embodiment.Set(w.Embodiment)
	tri.label.Set(w.Label)
	tri.state.Set(w.State)
	return tri
}

type Tristate struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	label      StringField
	state      IntegerField
}

func (tri *Tristate) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	tri.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &tri.embodiment},
			{key.FKey_Label, &tri.label},
			{key.FKey_State, &tri.state},
		}
	})
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
