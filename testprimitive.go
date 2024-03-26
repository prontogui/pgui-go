package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type TestPrimitive struct {
	Reserved

	BSide BSide
	B     field.Boolean
	S     field.String
	I     field.Integer
	SA    field.Strings1D
	A1D   field.Any1D
	A2D   field.Any2D
	BL    field.Blob
}

func (tp *TestPrimitive) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {
	tp.AttachField("B", &tp.B)
	tp.AttachField("S", &tp.S)
	tp.AttachField("I", &tp.I)
	tp.AttachField("SA", &tp.SA)
	tp.AttachField("A1D", &tp.A1D)
	tp.AttachField("A2D", &tp.A2D)
	tp.AttachField("BL", &tp.BL)

	// Prepare all the field for updates
	tp.Reserved.PrepareForUpdates(pkey, onset, &tp.BSide)
}
