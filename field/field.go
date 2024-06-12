// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package field

import "github.com/prontogui/golib/key"

type Field interface {
	PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) (isContainer bool)
	EgestValue() any
	IngestValue(value any) error
}
