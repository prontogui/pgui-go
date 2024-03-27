package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type ComplexPrimitive struct {
	Reserved

	BSide      BSide
	Issued     field.Boolean
	Embodiment field.String
	Status     field.Integer
	Choices    field.Strings1D
	ListItems  field.Any1D
	Rows       field.Any2D
	Data       field.Blob
}

func (tp *ComplexPrimitive) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {
	tp.AttachField("Issued", &tp.Issued)
	tp.AttachField("Embodiment", &tp.Embodiment)
	tp.AttachField("Status", &tp.Status)
	tp.AttachField("Choices", &tp.Choices)
	tp.AttachField("ListItems", &tp.ListItems)
	tp.AttachField("Rows", &tp.Rows)
	tp.AttachField("Data", &tp.Data)

	// Prepare all the field for updates
	tp.Reserved.PrepareForUpdates(pkey, onset, &tp.BSide)
}

type SimplePrimitive struct {
	Reserved

	BSide      BSide
	Issued     field.Boolean
	Embodiment field.String
	Status     field.Integer
}

func (tp *SimplePrimitive) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {
	tp.AttachField("Issued", &tp.Issued)
	tp.AttachField("Embodiment", &tp.Embodiment)
	tp.AttachField("Status", &tp.Status)

	// Prepare all the field for updates
	tp.Reserved.PrepareForUpdates(pkey, onset, &tp.BSide)
}
