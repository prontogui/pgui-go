package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_CommandAttachedFields(t *testing.T) {
	cmd := &Command{}
	cmd.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, cmd.Reserved, "Embodiment", "Label", "Status")
}

func Test_CommandMake(t *testing.T) {
	cmd := CommandWith{Embodiment: "raised-btn", Label: "Press Me", Status: 1}.Make()

	if cmd.Embodiment() != "raised-btn" {
		t.Error("Could not initialize Embodiment field.")
	}

	if cmd.Label() != "Press Me" {
		t.Error("Could not initialize Label field.")
	}

	if cmd.Status() != 1 {
		t.Error("Could not initialize Status field.")
	}
}

func Test_CommandFieldSetting(t *testing.T) {
	cmd := &Command{}
	cmd.PrepareForUpdates(key.NewPKey(), nil)

	cmd.SetEmbodiment("raised-btn")
	if cmd.Embodiment() != "raised-btn" {
		t.Error("Could not set Embodiment field.")
	}

	cmd.SetLabel("My label")
	if cmd.Label() != "My label" {
		t.Error("Could not set Label field.")
	}

	cmd.SetStatus(2)
	if cmd.Status() != 2 {
		t.Error("Could not set Status field.")
	}
}
