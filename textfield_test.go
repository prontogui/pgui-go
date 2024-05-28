package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_TextFieldAttach(t *testing.T) {
	txt := &TextField{}
	txt.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, txt.Reserved, "TextEntry", "Embodiment")
}

func Test_TextFieldMake(t *testing.T) {
	txt := TextFieldWith{
		TextEntry:  "This is a piece of text",
		Embodiment: "block",
	}.Make()

	if txt.TextEntry() != "This is a piece of text" {
		t.Error("Could not initialize TextEntry field.")
	}

	if txt.Embodiment() != "block" {
		t.Error("Could not initialize Embodiment field.")
	}
}

func Test_TextFieldFieldSetting(t *testing.T) {
	txt := &TextField{}
	txt.SetTextEntry("This is some nice content.")
	if txt.TextEntry() != "This is some nice content." {
		t.Error("Could not set Content field.")
	}

	txt.SetEmbodiment("block")
	if txt.Embodiment() != "block" {
		t.Error("Could not set Embodiment fields.")
	}
}
