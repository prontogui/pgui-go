// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package field

import (
	"github.com/prontogui/golib/key"
)

type Reserved struct {
	// PKey of this field's container primitive.
	pkey key.PKey // `cbor:"omitempty"`

	// FKey of this field.
	fkey key.FKey

	// The function to call to notify the field was updated.
	onset func(key.PKey, key.FKey, bool)

	// This field's pkey index relative to its container primitive (if this field contains primitives).
	// It is used when assigning new primitives to this field.
	fieldPKeyIndex int
}

func (f *Reserved) StashUpdateInfo(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) {
	f.fkey = fkey
	f.pkey = pkey
	f.onset = onset
	fieldPKeyIndex = fieldPKeyIndex
}

func (f *Reserved) OnSet(structural bool) {
	if f.onset != nil {
		f.onset(f.pkey, f.fkey, structural)
	}
}
