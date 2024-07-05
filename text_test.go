// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_TextAttach(t *testing.T) {
	txt := &Text{}
	txt.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, txt.PrimitiveBase, "Content", "Embodiment")
}

func Test_TextMake(t *testing.T) {
	txt := TextWith{
		Content:    "This is a piece of text",
		Embodiment: "block",
	}.Make()

	if txt.Content() != "This is a piece of text" {
		t.Error("Could not initialize Content field.")
	}

	if txt.Embodiment() != "block" {
		t.Error("Could not initialize Embodiment field.")
	}
}

func Test_TextFieldSetting(t *testing.T) {
	txt := &Text{}
	txt.SetContent("This is some nice content.")
	if txt.Content() != "This is some nice content." {
		t.Error("Could not set Content field.")
	}

	txt.SetEmbodiment("block")
	if txt.Embodiment() != "block" {
		t.Error("Could not set Embodiment fields.")
	}
}
