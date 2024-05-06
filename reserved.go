package golib

import (
	"errors"
	"fmt"

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
	bside  BSide
}

func (r *Reserved) B() *BSide {
	return &r.bside
}

func (r *Reserved) AttachField(fieldname string, field field.Field) {

	fkey := key.FKeyFor(fieldname)
	if fkey == key.INVALID_FIELDNAME {
		panic(fmt.Sprintf("Field name '%s' is not registered in key package.", fieldname))
	}

	r.fields = append(r.fields, FieldRef{fkey: fkey, field: field})
}

func (r *Reserved) GetChildPrimitive(index int) primitive.Interface {
	return nil
}

func (r *Reserved) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	r.bside.AttachFields(r)

	for _, f := range r.fields {
		f.field.PrepareForUpdates(f.fkey, pkey, onset)
	}
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
