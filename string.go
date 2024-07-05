// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type StringField struct {
	FieldBase
	s string
}

func (f *StringField) Get() string {
	return f.s
}

func (f *StringField) Set(s string) {
	f.s = s
	f.OnSet(false)
}

func (f *StringField) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) (isContainer bool) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset)
	return false
}

func (f *StringField) EgestValue() any {
	return f.s
}

func (f *StringField) IngestValue(value any) error {

	s, ok := value.(string)

	if !ok {
		return errors.New("unable to convert value (any) to field value")
	}

	f.s = s

	return nil
}
