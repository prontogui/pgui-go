// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package field

import (
	"errors"

	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

type Any1D struct {
	Reserved
	ary []primitive.Interface
}

func (f *Any1D) prepareDescendantsForUpdate() {

	fieldPkey := f.pkey.AddLevel(f.fieldPKeyIndex)

	for i, p := range f.ary {
		if f.onset == nil {
			p.PrepareForUpdates(key.EmptyPKey(), nil)
		} else {
			p.PrepareForUpdates(fieldPkey.AddLevel(i), f.onset)
		}
	}
}

func (f *Any1D) Get() []primitive.Interface {
	return f.ary
}

func (f *Any1D) Set(ary []primitive.Interface) {
	f.ary = ary
	f.prepareDescendantsForUpdate()
	f.OnSet(true)
}

func (f *Any1D) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) (isContainer bool) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset)
	f.prepareDescendantsForUpdate()
	return true
}

func (f *Any1D) EgestValue() any {

	ary := []any{}

	for _, v := range f.ary {
		ary = append(ary, v.EgestUpdate(true, nil))
	}

	return ary
}

func (f *Any1D) IngestValue(value any) error {

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
