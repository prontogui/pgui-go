// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

type FrameWith struct {
	Embodiment string
	Showing    bool
	FrameItems []primitive.Interface
}

func (w FrameWith) Make() *Frame {
	frame := &Frame{}
	frame.embodiment.Set(w.Embodiment)
	frame.showing.Set(w.Showing)
	frame.frameItems.Set(w.FrameItems)
	return frame
}

type Frame struct {
	// Mix-in the common guts for primitives
	Reserved

	embodiment field.String
	showing    field.Boolean
	frameItems field.Any1D
}

func (frame *Frame) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	frame.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &frame.embodiment},
			{key.FKey_FrameItems, &frame.frameItems},
			{key.FKey_Showing, &frame.showing},
		}
	})
}

// TODO:  generalize this code by handling inside primitive Reserved area.
func (frame *Frame) LocateNextDescendant(locator *key.PKeyLocator) primitive.Interface {
	if locator.NextIndex() != 0 {
		panic("cannot locate descendent using a pkey that we assumed was valid")
	}
	return frame.FrameItems()[locator.NextIndex()]
}

func (frame *Frame) FrameItems() []primitive.Interface {
	return frame.frameItems.Get()
}

func (frame *Frame) SetFrameItems(items []primitive.Interface) {
	frame.frameItems.Set(items)
}

func (frame *Frame) SetFrameItemsVA(items ...primitive.Interface) {
	frame.frameItems.Set(items)
}

func (frame *Frame) Showing() bool {
	return frame.showing.Get()
}

func (frame *Frame) SetShowing(showing bool) {
	frame.showing.Set(showing)
}
