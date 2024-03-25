package synchro

import (
	"errors"

	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

type testcommand struct {
	Label  field.String
	Issued field.Boolean
	Status field.Integer
}

func (cmd *testcommand) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	// Prepare all the field for updates
	cmd.Label.PrepareForUpdates(0, pkey, onset)
	cmd.Issued.PrepareForUpdates(1, pkey, onset)
	cmd.Status.PrepareForUpdates(2, pkey, onset)
}

func (cmd *testcommand) GetChildPrimitive(index int) primitive.Interface {
	return nil
}

func (cmd *testcommand) GetFieldValue(fkey key.FKey) any {
	switch fkey {
	case 0:
		return cmd.Label.Get()
	case 1:
		return cmd.Issued.Get()
	case 2:
		return cmd.Status.Get()
	}
	return nil
}

func (cmd *testcommand) IngestFieldUpdate(fieldname string, update any) error {
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
