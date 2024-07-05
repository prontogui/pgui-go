// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type BooleanField struct {
	FieldBase
	b bool
}

func (f *BooleanField) GetAsAny() any {
	return f.b
}

func (f *BooleanField) Get() bool {
	return f.b
}

func (f *BooleanField) Set(b bool) {
	f.b = b
	f.OnSet(false)
}

func (f *BooleanField) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) (isContainer bool) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset)
	return false
}

func (f *BooleanField) EgestValue() any {
	return f.b
}

func (f *BooleanField) IngestValue(value any) error {

	b, ok := value.(bool)

	if !ok {
		return errors.New("unable to convert value (any) to field value")
	}

	f.b = b
	return nil
}
