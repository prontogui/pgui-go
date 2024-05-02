package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_CommandAttachedFields(t *testing.T) {
	cmd := &Command{}
	cmd.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, cmd.Reserved, "Label", "Issued", "Status")
}

func Test_CommandMake(t *testing.T) {
	cmd := CommandWith{Label: "Press Me", Status: 1}.Make()

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

	cmd.SetLabel("My label")
	if cmd.Label() != "My label" {
		t.Error("Could not set Label field.")
	}

	cmd.SetStatus(2)
	if cmd.Status() != 2 {
		t.Error("Could not set Status field.")
	}

	cmd.IngestUpdate(map[any]any{"Issued": true})
	if !cmd.Issued() {
		t.Error("Could not get Issued field correctly.")
	}
}
