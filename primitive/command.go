package primitive

import (
	"github.com/prontogui/golib/field"
)

type Command struct {
	PReserved

	BSide BSide

	Label  field.String
	Issued field.Boolean
	Status field.Integer
}

func NewCommand() *Command {
	cmd := &Command{}

	closure := func(fieldno int8) {
		cmd.OnFieldSet(fieldno)
	}

	// Initialize PReserved stuff according to number of fields
	cmd.Init(3)

	// Assign field number and "on set" closure to each field
	// TODO:  handle BSide somehow
	cmd.BSide.OnSet(closure)
	cmd.Label.OnSet(0, closure)
	// And other fields...

	return cmd
}
