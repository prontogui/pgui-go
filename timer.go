// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

type TimerWith struct {
	Embodiment string
	PeriodMs   int
}

// Makes a new Timer with specified field values.
func (w TimerWith) Make() *Timer {
	cmd := &Timer{}
	cmd.embodiment.Set(w.Embodiment)
	cmd.periodMs.Set(w.PeriodMs)
	return cmd
}

type Timer struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	periodMs   IntegerField
}

func (tmr *Timer) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	tmr.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &tmr.embodiment},
			{key.FKey_PeriodMs, &tmr.periodMs},
		}
	})
}

func (tmr *Timer) Embodiment() string {
	return tmr.embodiment.Get()
}

func (tmr *Timer) SetEmbodiment(s string) {
	tmr.embodiment.Set(s)
}

func (tmr *Timer) PeriodMs() int {
	return tmr.periodMs.Get()
}

func (tmr *Timer) SetPeriodMs(i int) {
	tmr.periodMs.Set(i)
}
