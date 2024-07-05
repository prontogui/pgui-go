// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type AnyField struct {
	FieldBase
	p Primitive
}

func (f *AnyField) prepareDescendantForUpdate() {
	if f.p != nil {
		if f.onset == nil {
			f.p.PrepareForUpdates(key.EmptyPKey(), nil)
		} else {
			f.p.PrepareForUpdates(f.pkey.AddLevel(f.fieldPKeyIndex), f.onset)
		}
	}
}

func (f *AnyField) Get() Primitive {
	return f.p
}

func (f *AnyField) Set(p Primitive) {
	f.p = p
	f.prepareDescendantForUpdate()
	f.OnSet(true)
}

func (f *AnyField) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) (isContainer bool) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset)
	f.prepareDescendantForUpdate()
	return true
}

func (f *AnyField) EgestValue() any {
	if f.p != nil {
		return f.p.EgestUpdate(true, nil)
	} else {
		return nil
	}
}

func (f *AnyField) IngestValue(value any) error {

	m, ok := value.(map[any]any)
	if !ok {
		return errors.New("invalid update")
	}

	if f.p != nil {
		return f.p.IngestUpdate(m)
	}

	return nil
}
