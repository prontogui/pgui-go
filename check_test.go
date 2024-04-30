package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_CheckAttachedFields(t *testing.T) {
	check := &Check{}
	check.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, check.Reserved, "Label", "Checked")
}

func Test_CheckFieldSettings(t *testing.T) {

	check := &Check{}

	check.SetLabel("Option 1")

	if check.Label() != "Option 1" {
		t.Error("Could not set Label field.")
	}

	check.SetChecked(true)

	if !check.Checked() {
		t.Error("Could not set Checked field.")
	}
}
