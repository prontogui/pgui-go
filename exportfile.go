// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type ExportFileWith struct {
	Data       []byte
	Embodiment string
	Name       string
}

// Makes a new Command with specified field values.
func (w ExportFileWith) Make() *ExportFile {
	ef := &ExportFile{}
	ef.data.Set(w.Data)
	ef.embodiment.Set(w.Embodiment)
	ef.name.Set(w.Name)
	return ef
}

type ExportFile struct {
	// Mix-in the common guts for primitives
	Reserved

	data       field.Blob
	embodiment field.String
	exported   field.Boolean
	name       field.String
}

func (ef *ExportFile) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	ef.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Data, &ef.data},
			{key.FKey_Embodiment, &ef.embodiment},
			{key.FKey_Exported, &ef.exported},
			{key.FKey_Name, &ef.name},
		}
	})
}

func (ef *ExportFile) Data() []byte {
	return ef.data.Get()
}

func (ef *ExportFile) SetData(d []byte) {
	ef.data.Set(d)
}

func (ef *ExportFile) Embodiment() string {
	return ef.embodiment.Get()
}

func (ef *ExportFile) SetEmbodiment(s string) {
	ef.embodiment.Set(s)
}

func (ef *ExportFile) Exported() bool {
	return ef.exported.Get()
}

func (ef *ExportFile) SetExported(b bool) {
	ef.exported.Set(b)
}

func (ef *ExportFile) Name() string {
	return ef.name.Get()
}

func (ef *ExportFile) SetName(s string) {
	ef.name.Set(s)
}
