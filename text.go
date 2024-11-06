// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A text primitive displays text on the screen.
type TextWith struct {
	Content    string
	Embodiment string
	Tag        string
}

// Creates a new Text primitive using the supplief field assignments.
func (w TextWith) Make() *Text {
	text := &Text{}
	text.content.Set(w.Content)
	text.embodiment.Set(CanonizeEmbodiment(w.Embodiment))
	text.tag.Set(w.Tag)
	return text
}

// A text primitive displays text on the screen.
type Text struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	content    StringField
	embodiment StringField
	tag        StringField
}

// Create a new Text and assign its content.
func NewText(content string) *Text {
	return TextWith{Content: content}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (txt *Text) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	txt.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Content, &txt.content},
			{key.FKey_Embodiment, &txt.embodiment},
			{key.FKey_Tag, &txt.tag},
		}
	})
}

// Returns a string representation of this primitive:  the content.
// Implements of fmt:Stringer interface.
func (txt *Text) String() string {
	return txt.content.Get()
}

// Returns the text content to display.
func (txt *Text) Content() string {
	return txt.content.Get()
}

// Sets the text content to display.
func (txt *Text) SetContent(s string) *Text {
	txt.content.Set(s)
	return txt
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (txt *Text) Embodiment() string {
	return txt.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (txt *Text) SetEmbodiment(s string) *Text {
	txt.embodiment.Set(CanonizeEmbodiment(s))
	return txt
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Texts as Table cells.
func (txt *Text) Tag() string {
	return txt.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Texts as Table cells.
func (txt *Text) SetTag(s string) *Text {
	txt.tag.Set(s)
	return txt
}
