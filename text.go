package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type Text struct {
	// Mix-in the common guts for primitives (B-side fields, implements primitive interface, etc.)
	Reserved

	content field.String
}

func (txt *Text) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	txt.AttachField("Content", &txt.content)

	// Prepare all fields for updates
	txt.Reserved.PrepareForUpdates(pkey, onset)
}

func (txt *Text) Content() string {
	return txt.content.Get()
}

func (txt *Text) SetContent(s string) {
	txt.content.Set(s)
}
