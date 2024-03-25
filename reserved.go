package golib

import (
	"errors"

	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

const (
	// The maximum number of fields in any given primitive.  TODO:  check for accuracy of this in unit testing,
	// in case a primitive is updated or added without changing this number.
	MaxPrimitiveFields = 4
)

type FieldRef struct {
	fkey  key.FKey
	field field.Field
}

/*
Reserved fields for primitive updates.
*/
type Reserved struct {
	fields []FieldRef
	bside  primitive.Interface
}

func (r *Reserved) AttachField(fieldname string, field field.Field) {

	fkey := key.FKeyFor(fieldname)

	r.fields = append(r.fields, FieldRef{fkey: fkey, field: field})
}

func (r *Reserved) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, bside primitive.Interface) {

	r.bside = bside

	for _, f := range r.fields {
		f.field.PrepareForUpdates(f.fkey, pkey, onset)
	}
}

func (r *Reserved) GetChildPrimitive(index int) primitive.Interface {

	if index == 0 {
		return r.bside
	}
	return nil
}

func (r *Reserved) GetFieldValue(fkey key.FKey) any {

	for _, f := range r.fields {
		if f.fkey == fkey {
			return f.field.GetAsAny()
		}
	}
	return nil
}

func (r *Reserved) IngestFieldUpdate(fieldname string, update any) error {

	fkey := key.FKeyFor(fieldname)
	for _, f := range r.fields {
		if f.fkey == fkey {
			return f.field.IngestUpdate(update)
		}
	}
	return errors.New("field not found")
}
