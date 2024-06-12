// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"slices"
	"testing"

	"github.com/prontogui/golib/key"
)

func _areFieldsAttachedAlphabetically(res Reserved) bool {

	attachedOrder := []string{}

	for _, fr := range res.fields {
		fieldName := key.FieldnameFor(fr.fkey)
		attachedOrder = append(attachedOrder, fieldName)
	}

	return slices.IsSorted(attachedOrder)
}

func verifyAllFieldsAttached(t *testing.T, res Reserved, fields ...string) {

	verifyFieldAttached := func(fields ...string) {
		for _, field := range fields {
			if res.findField(key.FKeyFor(field)) == nil {
				t.Errorf("Field '%s' is not attached to primitive.", field)
			}
		}
	}

	verifyFieldAttached(fields...)

	if !_areFieldsAttachedAlphabetically(res) {
		t.Error("fields were not attached in alphabetical order")
	}
}
