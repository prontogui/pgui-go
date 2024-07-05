// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import "github.com/prontogui/golib/key"

type Primitive interface {
	PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction)
	LocateNextDescendant(locator *key.PKeyLocator) Primitive
	EgestUpdate(fullupdate bool, fkeys []key.FKey) map[any]any
	IngestUpdate(update map[any]any) error
}
