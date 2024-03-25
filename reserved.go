package golib

import (
	"github.com/prontogui/golib/primitive"
)

const (
	// The maximum number of fields in any given primitive.  TODO:  check for accuracy of this in unit testing,
	// in case a primitive is updated or added without changing this number.
	MaxPrimitiveFields = 4
)

/*
Reserved fields for primitive updates.
*/
type Reserved struct {
	BSide BSide
}

func (cmd *Reserved) GetChildPrimitive(index int) primitive.Interface {
	if index == 0 {
		return &cmd.BSide
	}
	return nil
}
