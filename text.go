// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type TextWith struct {
	Content    string
	Embodiment string
}

func (w TextWith) Make() *Text {
	text := &Text{}
	text.content.Set(w.Content)
	text.embodiment.Set(w.Embodiment)
	return text
}

type Text struct {
	// Mix-in the common guts for primitives
	Reserved

	content    field.String
	embodiment field.String
}

func (txt *Text) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	txt.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Content, &txt.content},
			{key.FKey_Embodiment, &txt.embodiment},
		}
	})
}

func (txt *Text) Content() string {
	return txt.content.Get()
}

func (txt *Text) SetContent(s string) {
	txt.content.Set(s)
}

func (txt *Text) Embodiment() string {
	return txt.embodiment.Get()
}

func (txt *Text) SetEmbodiment(s string) {
	txt.embodiment.Set(s)
}
