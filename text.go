package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type TextWith struct {
	Content string
}

func (w TextWith) Make() *Text {
	text := &Text{}
	text.content.Set(w.Content)
	return text
}

type Text struct {
	// Mix-in the common guts for primitives
	Reserved

	content field.String
}

func (txt *Text) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {
	txt.AttachField("Content", &txt.content, pkey, PKeyIndexDontCare, onset)
}

func (txt *Text) Content() string {
	return txt.content.Get()
}

func (txt *Text) SetContent(s string) {
	txt.content.Set(s)
}
