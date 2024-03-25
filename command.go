package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type Command struct {
	Reserved

	BSide  BSide
	Label  field.String
	Issued field.Boolean
	Status field.Integer
}

func (cmd *Command) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	cmd.AttachField("Label", &cmd.Label)
	cmd.AttachField("Issued", &cmd.Issued)
	cmd.AttachField("Status", &cmd.Status)

	// Prepare all the field for updates
	cmd.Reserved.PrepareForUpdates(pkey, onset, &cmd.BSide)
}
