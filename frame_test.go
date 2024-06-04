package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

func Test_FrameAttachedFields(t *testing.T) {
	frame := &Frame{}
	frame.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, frame.Reserved, "Embodiment", "Showing", "FrameItems")
}

func Test_FrameMake(t *testing.T) {
	frame := FrameWith{Showing: true, Embodiment: "full-view", FrameItems: []primitive.Interface{&Command{}, &Command{}}}.Make()

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

	frame.SetFrameItems([]primitive.Interface{&Command{}, &Command{}})

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

func Test_FrameGetChildPrimitive(t *testing.T) {

	frame := &Frame{}

	frame.SetFrameItems([]primitive.Interface{&Command{}, &Command{}})

	if frame.GetChildPrimitive(0) == nil {
		t.Fatal("GetChildPrimitve doesn't return a child for index = 0.")
	}

	if frame.GetChildPrimitive(1) == nil {
		t.Fatal("GetChildPrimitve doesn't return a child for index = 1.")
	}

	if frame.GetChildPrimitive(2) != nil {
		t.Fatal("GetChildPrimitve shouldn't return a child for index = 2.")
	}
}
