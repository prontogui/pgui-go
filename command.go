package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type CommandWith struct {
	Label  string
	Status int
}

// Makes a new Command with specified field values.
func (w CommandWith) Make() *Command {
	cmd := &Command{}
	cmd.label.Set(w.Label)
	cmd.status.Set(w.Status)
	return cmd
}

type Command struct {
	// Mix-in the common guts for primitives (B-side fields, implements primitive interface, etc.)
	Reserved

	label  field.String
	status field.Integer
	issued field.Event
}

func (cmd *Command) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	cmd.AttachField("Label", &cmd.label)
	cmd.AttachField("Issued", &cmd.issued)
	cmd.AttachField("Status", &cmd.status)

	// Prepare all fields for updates
	cmd.Reserved.PrepareForUpdates(pkey, onset)
}

func (cmd *Command) Label() string {
	return cmd.label.Get()
}

func (cmd *Command) SetLabel(s string) {
	cmd.label.Set(s)
}

func (cmd *Command) Status() int {
	return cmd.status.Get()
}

func (cmd *Command) SetStatus(i int) {
	cmd.status.Set(i)
}

func (cmd *Command) Issued() bool {
	return cmd.issued.Get()
}
