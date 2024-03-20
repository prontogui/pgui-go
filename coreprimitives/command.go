package coreprimitives

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type Command struct {
	Reserved

	BSide BSide

	Label  field.String
	Issued field.Boolean
	Status field.Integer
}

func (cmd *Command) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {
	// Prepare all the field for updates
	cmd.BSide.PrepareForUpdates(pkey, onset)
	cmd.Label.PrepareForUpdates("Label", pkey, onset)
	cmd.Issued.PrepareForUpdates("Issued", pkey, onset)
	cmd.Status.PrepareForUpdates("Status", pkey, onset)
}
