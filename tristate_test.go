package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_TristateAttachedFields(t *testing.T) {
	tri := &Tristate{}
	tri.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, tri.Reserved, "Label", "State")
}

func Test_TristateFieldSettings(t *testing.T) {

	tri := &Tristate{}

	tri.SetLabel("Yes, No, Undecided")

	if tri.Label() != "Yes, No, Undecided" {
		t.Error("Could not set Label field.")
	}

	tri.SetState(2)

	if tri.State() != 2 {
		t.Error("Could not set State field.")
	}
}
