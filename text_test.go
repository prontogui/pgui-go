package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_TextAttach(t *testing.T) {
	txt := &Text{}
	txt.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, txt.Reserved, "Content")
}

func Test_TextFieldSetting(t *testing.T) {
	txt := &Text{}
	txt.SetContent("This is some nice content.")
	if txt.Content() != "This is some nice content." {
		t.Error("Could not set Content field.")
	}
}
