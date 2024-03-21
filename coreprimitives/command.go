package coreprimitives

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
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

func (cmd *Command) GetChildPrimitive(index int) primitive.Interface {
	if index == 0 {
		return &cmd.BSide
	}
	return nil
}

func (cmd *Command) GetFieldValue(fieldname string) any {
	switch fieldname {
	case "Label":
		return cmd.Label.Get()
	case "Issued":
		return cmd.Issued.Get()
	case "Status":
		return cmd.Status.Get()
	}
	return nil
}
