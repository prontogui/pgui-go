// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type ImportFileWith struct {
	Data            []byte
	Embodiment      string
	Name            string
	ValidExtensions []string
}

// Makes a new Command with specified field values.
func (w ImportFileWith) Make() *ImportFile {
	ifile := &ImportFile{}
	ifile.data.Set(w.Data)
	ifile.embodiment.Set(w.Embodiment)
	ifile.name.Set(w.Name)
	ifile.validExtensions.Set(w.ValidExtensions)
	return ifile
}

type ImportFile struct {
	// Mix-in the common guts for primitives
	Reserved

	data            field.Blob
	embodiment      field.String
	imported        field.Boolean
	name            field.String
	validExtensions field.Strings1D
}

func (ifile *ImportFile) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	ifile.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Data, &ifile.data},
			{key.FKey_Embodiment, &ifile.embodiment},
			{key.FKey_Imported, &ifile.imported},
			{key.FKey_Name, &ifile.name},
			{key.FKey_ValidExtensions, &ifile.validExtensions},
		}
	})
}

func (ifile *ImportFile) Data() []byte {
	return ifile.data.Get()
}

func (ifile *ImportFile) SetData(d []byte) {
	ifile.data.Set(d)
}

func (ifile *ImportFile) Embodiment() string {
	return ifile.embodiment.Get()
}

func (ifile *ImportFile) SetEmbodiment(s string) {
	ifile.embodiment.Set(s)
}

func (ifile *ImportFile) Imported() bool {
	return ifile.imported.Get()
}

func (ifile *ImportFile) SetImported(b bool) {
	ifile.imported.Set(b)
}

func (ifile *ImportFile) Name() string {
	return ifile.name.Get()
}

func (ifile *ImportFile) SetName(s string) {
	ifile.name.Set(s)
}

func (ifile *ImportFile) ValidExtensions() []string {
	return ifile.validExtensions.Get()
}

func (ifile *ImportFile) SetValidExtensions(sa []string) {
	ifile.validExtensions.Set(sa)
}
