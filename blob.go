// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type BlobField struct {
	FieldBase
	blob []byte
}

func (f *BlobField) Get() []byte {
	return f.blob
}

func (f *BlobField) Set(blob []byte) {
	f.blob = blob
	f.OnSet(false)
}

func (f *BlobField) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) (isContainer bool) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset)
	return false
}

func (f *BlobField) EgestValue() any {
	return f.blob
}

func (f *BlobField) IngestValue(value any) error {

	bytes, ok := value.([]uint8)
	if !ok {
		return errors.New("ingested value type not supported for Blob")
	}

	f.blob = bytes
	return nil
}
