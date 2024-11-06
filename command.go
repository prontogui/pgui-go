// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A command is used to handle momentary requests by the user such that, when the command is issued,
// the service does something useful.  It is often rendered as a button with clear boundaries that
// suggest it can be clicked.
type CommandWith struct {
	Embodiment string
	Label      string
	Status     int
	Tag        string
}

// Makes a new Command with specified field values.
func (w CommandWith) Make() *Command {
	cmd := &Command{}

	// Must initialize the CommandIssued field
	cmd.commandIssued.TimestampProvider = getEventTimestamp

	// Set fields
	cmd.embodiment.Set(CanonizeEmbodiment(w.Embodiment))
	cmd.label.Set(w.Label)
	cmd.status.Set(w.Status)
	cmd.tag.Set(w.Tag)

	return cmd
}

// A command is used to handle momentary requests by the user such that, when the command is issued,
// the service does something useful.  It is often rendered as a button with clear boundaries that
// suggest it can be clicked.
type Command struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	commandIssued EventField
	embodiment    StringField
	label         StringField
	status        IntegerField
	tag           StringField
}

// Creates a new command and assigns a label.
func NewCommand(label string) *Command {
	cmd := CommandWith{Label: label}.Make()
	return cmd
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (cmd *Command) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	cmd.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_CommandIssued, &cmd.commandIssued},
			{key.FKey_Embodiment, &cmd.embodiment},
			{key.FKey_Label, &cmd.label},
			{key.FKey_Status, &cmd.status},
			{key.FKey_Tag, &cmd.tag},
		}
	})
}

// Returns a string representation of this primitive:  the label.
// Implements of fmt:Stringer interface.
func (cmd *Command) String() string {
	return cmd.label.Get()
}

// Returns true if the command was issued during the current Wait cycle.
func (cmd *Command) Issued() bool {
	return cmd.commandIssued.Issued()
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (cmd *Command) Embodiment() string {
	return cmd.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (cmd *Command) SetEmbodiment(s string) *Command {
	cmd.embodiment.Set(CanonizeEmbodiment(s))
	return cmd
}

// Returns the label to display in the command.
func (cmd *Command) Label() string {
	return cmd.label.Get()
}

// Sets the label to display in the command.
func (cmd *Command) SetLabel(s string) *Command {
	cmd.label.Set(s)
	return cmd
}

// Returns the status of the command:  0 = Command Normal, 1 = Command Disabled, 2 = Command Hidden.
func (cmd *Command) Status() int {
	return cmd.status.Get()
}

// Sets the status of the command:  0 = Command Normal, 1 = Command Disabled, 2 = Command Hidden.
func (cmd *Command) SetStatus(i int) *Command {
	cmd.status.Set(i)
	return cmd
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Commands as Table cells.
func (cmd *Command) Tag() string {
	return cmd.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Commands as Table cells.
func (cmd *Command) SetTag(s string) *Command {
	cmd.tag.Set(s)
	return cmd
}

// Returns the visibility of the command.  This is derived from the Status field.
func (cmd *Command) Visible() bool {
	return cmd.status.Get() != 2
}

// Sets the visibility of the command.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (cmd *Command) SetVisible(visible bool) *Command {
	if visible {
		cmd.status.Set(0)
	} else {
		cmd.status.Set(2)
	}
	return cmd
}

// Returns the enabled status of the command.  This is derived from the Status field.
func (cmd *Command) Enabled() bool {
	return cmd.status.Get() == 0
}

// Sets the enabled status of the command.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (cmd *Command) SetEnabled(enabled bool) *Command {
	if enabled {
		cmd.status.Set(0)
	} else {
		cmd.status.Set(1)
	}
	return cmd
}
