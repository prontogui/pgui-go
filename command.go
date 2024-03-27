package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type Command struct {
	Reserved

	BSide  BSide
	label  field.String
	issued field.Boolean
	status field.Integer
}

func (cmd *Command) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	cmd.AttachField("Label", &cmd.label)
	cmd.AttachField("Issued", &cmd.issued)
	cmd.AttachField("Status", &cmd.status)

	// Prepare all the field for updates
	cmd.Reserved.PrepareForUpdates(pkey, onset, &cmd.BSide)
}

func (cmd *Command) Label() string {
	return cmd.label.Get()
}

func (cmd *Command) SetLabel(s string) {
	cmd.label.Set(s)
}

func (cmd *Command) Issued() bool {
	return cmd.issued.Get()
}

func (cmd *Command) SetIssued(b bool) {
	cmd.issued.Set(b)
}

func (cmd *Command) Status() int {
	return cmd.status.Get()
}

func (cmd *Command) SetStatus(i int) {
	cmd.status.Set(i)
}
