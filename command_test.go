package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_Attach(t *testing.T) {
	cmd := &Command{}
	cmd.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, cmd.Reserved, "Label", "Issued", "Status")
	verifyBsideFieldsAttached(t, &cmd.BSide, cmd.Reserved)
}
