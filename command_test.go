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

func Test_CommandFieldSetting(t *testing.T) {
	cmd := &Command{}
	cmd.SetIssued(true)
	if !cmd.Issued() {
		t.Error("Could not set Issued field.")
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
