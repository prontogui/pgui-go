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

func Test_TristateMake(t *testing.T) {
	tri := TristateWith{Label: "Maybe, Yes, or No", State: 1}.Make()

	if tri.Label() != "Maybe, Yes, or No" {
		t.Error("Could not initialize Label field.")
	}
	if tri.State() != 1 {
		t.Error("Could not initialize State field.")
	}
}

func Test_TristateFieldSettings(t *testing.T) {

	tri := &Tristate{}
	tri.PrepareForUpdates(key.NewPKey(), nil)

	tri.SetLabel("Yes, No, Undecided")

	if tri.Label() != "Yes, No, Undecided" {
		t.Error("Could not set Label field.")
	}

	tri.SetState(2)

	if tri.State() != 2 {
		t.Error("Could not set State field.")
	}
}
