// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// An entry field that allows the user to enter text.
type TextFieldWith struct {
	Embodiment string
	Tag        string
	TextEntry  string
}

// Creates a new TextField using the supplied field assignments.
func (w TextFieldWith) Make() *TextField {
	textField := &TextField{}
	textField.embodiment.Set(CanonizeEmbodiment(w.Embodiment))
	textField.textEntry.Set(w.TextEntry)
	textField.tag.Set(w.Tag)
	return textField
}

// An entry field that allows the user to enter text.
type TextField struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	tag        StringField
	textEntry  StringField
}

// Create a new TextField with initial text.
func NewTextField(textEntry string) *TextField {
	return TextFieldWith{TextEntry: textEntry}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (txt *TextField) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	txt.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &txt.embodiment},
			{key.FKey_Tag, &txt.tag},
			{key.FKey_TextEntry, &txt.textEntry},
		}
	})
}

// Returns a string representation of this primitive:  the text entry.
// Implements of fmt:Stringer interface.
func (txt *TextField) String() string {
	return txt.textEntry.Get()
}

// Returns the text entered by the user.
func (txt *TextField) TextEntry() string {
	return txt.textEntry.Get()
}

// Sets the text entered by the user.
func (txt *TextField) SetTextEntry(s string) *TextField {
	txt.textEntry.Set(s)
	return txt
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (txt *TextField) Embodiment() string {
	return txt.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (txt *TextField) SetEmbodiment(s string) *TextField {
	txt.embodiment.Set(CanonizeEmbodiment(s))
	return txt
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using TextFields as Table cells.
func (txt *TextField) Tag() string {
	return txt.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using TextFields as Table cells.
func (txt *TextField) SetTag(s string) *TextField {
	txt.tag.Set(s)
	return txt
}
