// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type Any1DField struct {
	FieldBase
	ary []Primitive
}

func (f *Any1DField) prepareDescendantsForUpdate() {

	fieldPkey := f.pkey.AddLevel(f.fieldPKeyIndex)

	for i, p := range f.ary {
		if f.onset == nil {
			p.PrepareForUpdates(key.EmptyPKey(), nil)
		} else {
			p.PrepareForUpdates(fieldPkey.AddLevel(i), f.onset)
		}
	}
}

func (f *Any1DField) Get() []Primitive {
	return f.ary
}

func (f *Any1DField) Set(ary []Primitive) {
	f.ary = ary
	f.prepareDescendantsForUpdate()
	f.OnSet(true)
}

func (f *Any1DField) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) (isContainer bool) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset)
	f.prepareDescendantsForUpdate()
	return true
}

func (f *Any1DField) EgestValue() any {

	ary := []any{}

	for _, v := range f.ary {
		ary = append(ary, v.EgestUpdate(true, nil))
	}

	return ary
}

func (f *Any1DField) IngestValue(value any) error {

	l, ok := value.([]any)
	if !ok {
		return errors.New("invalid update")
	}

	if len(l) != len(f.ary) {
		return errors.New("number of primitives in update does not equal existing primitives")
	}

	for i, v := range l {
		m, ok := v.(map[any]any)
		if !ok {
			return errors.New("invalid update")
		}

		err := f.ary[i].IngestUpdate(m)
		if err != nil {
			return err
		}
	}

	return nil
}
