// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_ImageAttachedFields(t *testing.T) {
	image := &Image{}
	image.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, image.PrimitiveBase, "Embodiment", "Image")
}

func Test_ImageMake1(t *testing.T) {
	image := ImageWith{Embodiment: "black-white", Image: []byte{0, 1, 2}}.Make()

	if image.Embodiment() != "black-white" {
		t.Error("Could not initialize Embodiment field.")
	}

	if len(image.Image()) != 3 {
		t.Error("Could not initialize Image field.")
	}
}

func Test_ImageMake2(t *testing.T) {
	image := ImageWith{Embodiment: "black-white", FromFile: "gopher.png"}.Make()

	if image.Embodiment() != "black-white" {
		t.Error("Could not initialize Embodiment field.")
	}

	if len(image.Image()) == 0 {
		t.Error("Could not initialize Image field using FromFile.")
	}
}

func Test_ImageFieldSetting(t *testing.T) {
	image := &Image{}
	image.PrepareForUpdates(key.NewPKey(), nil)

	image.SetEmbodiment("black-white")
	if image.Embodiment() != "black-white" {
		t.Error("Could not set Embodiment field.")
	}

	image.SetImage([]byte{0, 1, 2})
	if len(image.Image()) != 3 {
		t.Error("Could not set Image field.")
	}
}
