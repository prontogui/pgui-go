package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

type BSide struct {
	Reserved

	// The row index (0-based) where the primitive resides in a container (usually a table).
	// This is assigned automatically by the container primitive, or in the case of a
	// template row, by a mutation from the App after an event occured.
	Row field.Integer

	// The colunn index (0-based) where the primitive resides in a container (usually
	// a table).  This is assigned automatically by the container primitive, or in the case
	// of a template row, by a mutation from the App after an event occured.
	Col field.Integer

	// A JSON string specifying the embodiment to use for this primitive.
	// This can be assigned explicitly or by a mutation from the App.
	Embodiment field.String
}

func (bs *BSide) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {
	bs.AttachField("Row", &bs.Row)

}

func (bs *BSide) GetChildPrimitive(index int) primitive.Interface {
	return nil
}
