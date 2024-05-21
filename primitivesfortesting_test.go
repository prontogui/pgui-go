package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type ComplexPrimitive struct {
	Reserved

	BSide     BSide
	Issued    field.Boolean
	Status    field.Integer
	Choices   field.Strings1D
	ListItems field.Any1D
	Rows      field.Any2D
	Data      field.Blob
}

func (tp *ComplexPrimitive) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {
	tp.AttachField("Choices", &tp.Choices, pkey, PKeyIndexDontCare, onset)
	tp.AttachField("Data", &tp.Data, pkey, PKeyIndexDontCare, onset)
	tp.AttachField("Issued", &tp.Issued, pkey, PKeyIndexDontCare, onset)
	tp.AttachField("ListItems", &tp.ListItems, pkey, PKeyIndex_0, onset)
	tp.AttachField("Rows", &tp.Rows, pkey, PKeyIndex_1, onset)
	tp.AttachField("Status", &tp.Status, pkey, PKeyIndexDontCare, onset)
}

type SimplePrimitive struct {
	Reserved

	BSide  BSide
	Issued field.Boolean
	Label  field.String
	Status field.Integer
}

func (tp *SimplePrimitive) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {
	tp.AttachField("Issued", &tp.Issued, pkey, PKeyIndexDontCare, onset)
	tp.AttachField("Label", &tp.Label, pkey, PKeyIndexDontCare, onset)
	tp.AttachField("Status", &tp.Status, pkey, PKeyIndexDontCare, onset)
}
