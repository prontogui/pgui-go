// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package field

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type Integer struct {
	Reserved
	i int
}

func (f *Integer) Get() int {
	return f.i
}

func (f *Integer) Set(i int) {
	f.i = i
	f.OnSet(false)
}

func (f *Integer) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) (isContainer bool) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset)
	return false
}

func (f *Integer) EgestValue() any {
	return f.i
}

func (f *Integer) IngestValue(value any) error {

	// Unfortunately, CBOR encodes different sizes of integers based on optimum space usage.  It's not deterministic
	// what we are converting from.  So we have to test each case until a successful conversion happens.

	ui64, ok := value.(uint64)
	if ok {
		f.i = int(ui64)
		return nil
	}

	i64, ok := value.(int64)
	if ok {
		f.i = int(i64)
		return nil
	}

	i, ok := value.(int)
	if ok {
		f.i = i
		return nil
	}

	ui, ok := value.(uint)
	if ok {
		f.i = int(ui)
		return nil
	}

	ui32, ok := value.(uint32)
	if ok {
		f.i = int(ui32)
		return nil
	}

	i32, ok := value.(int32)
	if ok {
		f.i = int(i32)
		return nil
	}

	ui16, ok := value.(uint16)
	if ok {
		f.i = int(ui16)
		return nil
	}

	i16, ok := value.(int16)
	if ok {
		f.i = int(i16)
		return nil
	}

	ui8, ok := value.(uint8)
	if ok {
		f.i = int(ui8)
		return nil
	}

	i8, ok := value.(int8)
	if ok {
		f.i = int(i8)
		return nil
	}

	return errors.New("unable to convert value (any) to field value")
}
