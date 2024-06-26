// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package field

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type Blob struct {
	Reserved
	blob []byte
}

func (f *Blob) Get() []byte {
	return f.blob
}

func (f *Blob) Set(blob []byte) {
	f.blob = blob
	f.OnSet(false)
}

func (f *Blob) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) (isContainer bool) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset)
	return false
}

func (f *Blob) EgestValue() any {
	return f.blob
}

func (f *Blob) IngestValue(value any) error {

	bytes, ok := value.([]uint8)
	if !ok {
		return errors.New("ingested value type not supported for Blob")
	}

	f.blob = bytes
	return nil
}
