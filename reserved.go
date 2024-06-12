// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"errors"

	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

// Reserved fields for primitive updates.
type Reserved struct {
	fields []FieldRef
}

type FieldRef struct {
	// The field's key
	fkey key.FKey

	// Reference to the field itself
	field field.Field
}

func (r *Reserved) InternalPrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, getFields func() []FieldRef) {

	// Attach fields (if not done already)
	if len(r.fields) == 0 {
		r.fields = getFields()
	}

	// Prepare each field for updates
	fieldPKeyIndex := 0
	for _, f := range r.fields {
		if f.field.PrepareForUpdates(f.fkey, pkey, fieldPKeyIndex, onset) {
			fieldPKeyIndex = fieldPKeyIndex + 1
		}
	}
}

func (r *Reserved) LocateNextDescendant(locator *key.PKeyLocator) primitive.Interface {
	return nil
}

func (r *Reserved) findField(fkey key.FKey) field.Field {

	var found field.Field
	for _, f := range r.fields {
		if f.fkey == fkey {
			found = f.field
			break
		}
	}
	return found
}

func (r *Reserved) EgestUpdate(fullupdate bool, fkeys []key.FKey) map[any]any {

	update := map[any]any{}

	if fullupdate {
		for _, v := range r.fields {
			fieldvalue := v.field.EgestValue()

			if fieldvalue != nil {
				update[key.FieldnameFor(v.fkey)] = fieldvalue
			}
		}
	} else {
		for _, fkey := range fkeys {

			field := r.findField(fkey)
			if field == nil {
				panic("field not found in primitive")
			}

			fieldvalue := field.EgestValue()

			if fieldvalue != nil {
				update[key.FieldnameFor(fkey)] = fieldvalue
			}
		}
	}

	return update
}

func (r *Reserved) IngestUpdate(update map[any]any) error {

	for k, v := range update {
		var ok bool

		ks, ok := k.(string)
		if !ok {
			return errors.New("invalid key type.  Expecting a string")
		}

		fkey := key.FKeyFor(ks)
		if fkey == key.INVALID_FIELDNAME {
			return errors.New("invalid field name")
		}

		var found field.Field
		for _, f := range r.fields {
			if f.fkey == fkey {
				found = f.field
				break
			}
		}

		if found == nil {
			return errors.New("no matching field name in primitive")
		}

		err := found.IngestValue(v)
		if err != nil {
			return err
		}
	}

	return nil
}
