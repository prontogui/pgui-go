package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type ComplexPrimitive struct {
	Reserved

	Issued    field.Boolean
	Status    field.Integer
	Choices   field.Strings1D
	ListItems field.Any1D
	Rows      field.Any2D
	Data      field.Blob
}

func (tp *ComplexPrimitive) GetFieldRefs() []FieldRef {
	return []FieldRef{
		{key.FKey_Choices, &tp.Choices},
		{key.FKey_Data, &tp.Data},
		{key.FKey_Issued, &tp.Issued},
		{key.FKey_ListItems, &tp.ListItems},
		{key.FKey_Rows, &tp.Rows},
		{key.FKey_Status, &tp.Status},
	}
}

type SimplePrimitive struct {
	Reserved

	Issued field.Boolean
	Label  field.String
	Status field.Integer
}

func (tp *SimplePrimitive) GetFieldRefs() []FieldRef {
	return []FieldRef{
		{key.FKey_Issued, &tp.Issued},
		{key.FKey_Label, &tp.Label},
		{key.FKey_Status, &tp.Status},
	}
}
