// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A file represents a blob of data that can be exported from the server side and
// stored to a file on the app side.  The perspective of "export" is centered around
// the server software.  This seems to be a little clearer than using Download/Upload
// terminology.
type ExportFileWith struct {
	Data       []byte
	Embodiment string
	Name       string
	Tag        string
}

// Makes a new Command with specified field values.
func (w ExportFileWith) Make() *ExportFile {
	ef := &ExportFile{}
	ef.data.Set(w.Data)
	ef.embodiment.Set(CanonizeEmbodiment(w.Embodiment))
	ef.name.Set(w.Name)
	ef.tag.Set(w.Tag)
	return ef
}

// A file represents a blob of data that can be exported from the server side and
// stored to a file on the app side.  The perspective of "export" is centered around
// the server software.  This seems to be a little clearer than using Download/Upload
// terminology.
type ExportFile struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	data       BlobField
	embodiment StringField
	exported   BooleanField
	name       StringField
	tag        StringField
}

// Creates a new ExportFile.
func NewExportFile() *ExportFile {
	return ExportFileWith{}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (ef *ExportFile) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	ef.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Data, &ef.data},
			{key.FKey_Embodiment, &ef.embodiment},
			{key.FKey_Exported, &ef.exported},
			{key.FKey_Name, &ef.name},
			{key.FKey_Tag, &ef.tag},
		}
	})
}

// Returns the blob of data representing the binary contents of the file.  Note:  this
// data could be empty and yet represent a valid, albeit empty, file for export.
func (ef *ExportFile) Data() []byte {
	return ef.data.Get()
}

// Sets the blob of data representing the binary contents of the file to export.
func (ef *ExportFile) SetData(d []byte) *ExportFile {
	ef.data.Set(d)
	return ef
}

// Clears the exported data and the exported flag.
func (ef *ExportFile) Reset() {
	ef.data.Set([]byte{})
	ef.exported.Set(false)
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (ef *ExportFile) Embodiment() string {
	return ef.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (ef *ExportFile) SetEmbodiment(s string) *ExportFile {
	ef.embodiment.Set(CanonizeEmbodiment(s))
	return ef
}

// Returns true when the file has been exported (stored to a file) by the app.
// This field is normally only updated by the app.
func (ef *ExportFile) Exported() bool {
	return ef.exported.Get()
}

// Sets whether or not the file has been exported (stored to a file) by the app.
// This field is normally only updated by the app.
func (ef *ExportFile) SetExported(b bool) *ExportFile {
	ef.exported.Set(b)
	return ef
}

// Returns the suggested file name (including its extension separated by a period) to save the file as.
func (ef *ExportFile) Name() string {
	return ef.name.Get()
}

// Sets the suggested file name (including its extension separated by a period) to save the file as.
func (ef *ExportFile) SetName(s string) *ExportFile {
	ef.name.Set(s)
	return ef
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using ExportFiles as Table cells.
func (ef *ExportFile) Tag() string {
	return ef.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using ExportFiles as Table cells.
func (ef *ExportFile) SetTag(s string) *ExportFile {
	ef.tag.Set(s)
	return ef
}
