package golib

import (
	"errors"

	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type Command struct {
	Reserved

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

func (cmd *Command) IngestFieldUpdate(fieldname string, update any) error {

	switch fieldname {
	case "Label":
		return cmd.Label.IngestUpdate(update)
	case "Issued":
		return cmd.Issued.IngestUpdate(update)
	case "Status":
		return cmd.Status.IngestUpdate(update)
	}
	return errors.New("field not found")
}
