// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A timer is an invisible primitive that fires an event, triggering an update
// to the server.  This is useful for low-precision GUI updates that originate
// on the server side.  An example is updating "live" readings from a running
// process on the server.
type TimerWith struct {
	Embodiment string
	PeriodMs   int
	Tag        string
}

// Makes a new Timer with specified field values.
func (w TimerWith) Make() *Timer {
	timer := &Timer{}
	timer.embodiment.Set(CanonizeEmbodiment(w.Embodiment))
	timer.periodMs.Set(w.PeriodMs)
	timer.tag.Set(w.Tag)
	return timer
}

// A timer is an invisible primitive that fires an event, triggering an update
// to the server.  This is useful for low-precision GUI updates that originate
// on the server side.  An example is updating "live" readings from a running
// process on the server.
type Timer struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	periodMs   IntegerField
	tag        StringField
}

// Create a new Timer with period in milliseconds.
func NewTimer(periodMs int) *Timer {
	return TimerWith{PeriodMs: periodMs}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (tmr *Timer) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	tmr.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &tmr.embodiment},
			{key.FKey_PeriodMs, &tmr.periodMs},
			{key.FKey_Tag, &tmr.tag},
		}
	})
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (tmr *Timer) Embodiment() string {
	return tmr.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (tmr *Timer) SetEmbodiment(s string) *Timer {
	tmr.embodiment.Set(CanonizeEmbodiment(s))
	return tmr
}

// Returns the time period in milliseconds after which the timer fires an event.  If the
// period is -1 (or any negative number) then the timer is disabled.  A period of 0
// will cause the timer to fire immediately after the primitive is updated.
func (tmr *Timer) PeriodMs() int {
	return tmr.periodMs.Get()
}

// Sets the time period in milliseconds after which the timer fires an event.  If the
// period is -1 (or any negative number) then the timer is disabled.  A period of 0
// will cause the timer to fire immediately after the primitive is updated.
func (tmr *Timer) SetPeriodMs(i int) *Timer {
	tmr.periodMs.Set(i)
	return tmr
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Timers inside containers.
func (tmr *Timer) Tag() string {
	return tmr.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Timers inside containers.
func (tmr *Timer) SetTag(s string) *Timer {
	tmr.tag.Set(s)
	return tmr
}
