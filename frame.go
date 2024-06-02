package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

type FrameWith struct {
	FrameItems []primitive.Interface
}

func (w FrameWith) Make() *Frame {
	frame := &Frame{}
	frame.frameItems.Set(w.FrameItems)
	return frame
}

type Frame struct {
	// Mix-in the common guts for primitives (B-side fields, implements primitive interface, etc.)
	Reserved

	embodiment field.String
	frameItems field.Any1D
}

func (frame *Frame) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	frame.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &frame.embodiment},
			{key.FKey_FrameItems, &frame.frameItems},
		}
	})
}

func (frame *Frame) GetChildPrimitive(index int) primitive.Interface {
	frameItems := frame.frameItems.Get()

	if index < len(frameItems) {
		return frameItems[index]
	}

	return nil
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
