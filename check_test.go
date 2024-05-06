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

func Test_CheckMake(t *testing.T) {
	check := CheckWith{Label: "Option", Checked: true}.Make()

	if check.Label() != "Option" {
		t.Error("Could not initialize Label field.")
	}

	if !check.Checked() {
		t.Error("Could not initialize Checked field.")
	}
}

func Test_CheckFieldSettings(t *testing.T) {

	check := &Check{}
	check.PrepareForUpdates(key.NewPKey(), nil)

	check.SetLabel("Option 1")

	if check.Label() != "Option 1" {
		t.Error("Could not set Label field.")
	}

	check.SetChecked(true)

	if !check.Checked() {
		t.Error("Could not set Checked field.")
	}
}
