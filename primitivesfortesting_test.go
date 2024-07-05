// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

type ComplexPrimitive struct {
	PrimitiveBase

	Issued    BooleanField
	Status    IntegerField
	Choices   Strings1DField
	ListItems Any1DField
	Rows      Any2DField
	Data      BlobField
}

func (tp *ComplexPrimitive) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	tp.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Choices, &tp.Choices},
			{key.FKey_Data, &tp.Data},
			{key.FKey_Issued, &tp.Issued},
			{key.FKey_ListItems, &tp.ListItems},
			{key.FKey_Rows, &tp.Rows},
			{key.FKey_Status, &tp.Status},
		}
	})
}

type SimplePrimitive struct {
	PrimitiveBase

	Issued BooleanField
	Label  StringField
	Status IntegerField
}

func (tp *SimplePrimitive) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	tp.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Issued, &tp.Issued},
			{key.FKey_Label, &tp.Label},
			{key.FKey_Status, &tp.Status},
		}
	})
}
