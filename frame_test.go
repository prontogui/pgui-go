// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_FrameAttachedFields(t *testing.T) {
	frame := &Frame{}
	frame.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, frame.PrimitiveBase, "Embodiment", "Showing", "FrameItems")
}

func Test_FrameMake(t *testing.T) {
	frame := FrameWith{Showing: true, Embodiment: "full-view", FrameItems: []Primitive{&Command{}, &Command{}}}.Make()

	if !frame.showing.Get() {
		t.Error("'Showing' field was not initialized properly")
	}

	if frame.embodiment.Get() != "full-view" {
		t.Error("'Embodiment' field not initialized properly")
	}

	if len(frame.FrameItems()) != 2 {
		t.Error("'FrameItems' field was not initialized correctly")
	}
}

func Test_FrameFieldSettings(t *testing.T) {

	frame := &Frame{}

	frame.SetFrameItems([]Primitive{&Command{}, &Command{}})

	frameGet := frame.FrameItems()

	if len(frameGet) != 2 {
		t.Errorf("FrameItems() returned %d items.  Expecting 2 items.", len(frameGet))
	}

	_, ok1 := frameGet[0].(*Command)
	if !ok1 {
		t.Error("First frame item is not a Command primitive.")
	}
	_, ok2 := frameGet[1].(*Command)
	if !ok2 {
		t.Error("Second frame item is not a Command primitive.")
	}

	frame.SetFrameItemsVA(&Text{}, &Text{})

	frameGet = frame.FrameItems()

	if len(frameGet) != 2 {
		t.Errorf("GroupItems() returned %d items after calling variadic setter.  Expecting 2 items.", len(frameGet))
	}

	_, ok1 = frameGet[0].(*Text)
	if !ok1 {
		t.Error("First group is not a Text primitive.")
	}
	_, ok2 = frameGet[1].(*Text)
	if !ok2 {
		t.Error("Second group is not a Text primitive.")
	}

}

func Test_FrameLocateChildPrimitive(t *testing.T) {

	cmd1 := CommandWith{Label: "a"}.Make()
	cmd2 := CommandWith{Label: "b"}.Make()

	grp := FrameWith{FrameItems: []Primitive{cmd1, cmd2}}.Make()

	locate := func(pkey key.PKey) *Command {
		locator := key.NewPKeyLocator(pkey)
		return grp.LocateNextDescendant(locator).(*Command)
	}

	if locate(key.NewPKey(0, 0)).Label() != "a" {
		t.Fatal("LocateNextDescendant doesn't return a child for pkey 0.")
	}

	if locate(key.NewPKey(0, 1)).Label() != "b" {
		t.Fatal("LocateNextDescendant doesn't return a child for pkey 1.")
	}
}
