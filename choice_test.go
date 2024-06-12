// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"reflect"
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_ChoiceAttachedFields(t *testing.T) {
	cmd := &Choice{}
	cmd.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, cmd.Reserved, "Choice", "Choices", "Embodiment")
}

func Test_ChoiceMake(t *testing.T) {
	choice := ChoiceWith{Choice: "Apple", Choices: []string{"Apple", "Orange"}, Embodiment: "checkmark"}.Make()

	if choice.Choice() != "Apple" {
		t.Error("Could not initialize Choice field.")
	}

	if !reflect.DeepEqual(choice.Choices(), []string{"Apple", "Orange"}) {
		t.Error("Could not initialize Choices field.")
	}

	if choice.Embodiment() != "checkmark" {
		t.Error("Could not initialize Embodiment field.")
	}
}

func Test_ChoiceFieldSettings(t *testing.T) {
	choice := &Choice{}
	choice.PrepareForUpdates(key.NewPKey(), nil)

	choice.SetChoice("Big Fish")
	if choice.Choice() != "Big Fish" {
		t.Error("Could not set Choice field.")
	}

	choice.SetChoices([]string{"mary", "john", "paul"})
	if !reflect.DeepEqual(choice.Choices(), []string{"mary", "john", "paul"}) {
		t.Error("Could not set Choices field.")
	}

	choice.SetChoicesVA("nancy", "tom", "bob")
	if !reflect.DeepEqual(choice.Choices(), []string{"nancy", "tom", "bob"}) {
		t.Error("Could not set Choices field using variadic arguments")
	}

	choice.SetEmbodiment("checkmark")
	if choice.Embodiment() != "checkmark" {
		t.Error("Could not set Embodiment field")
	}
}
