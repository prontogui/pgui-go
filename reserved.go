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

func (r *Reserved) IngestUpdate(update map[any]any) (int, error) {

	reason := 0

	for k, v := range update {
		var ok bool

		ks, ok := k.(string)
		if !ok {
			return 0, errors.New("invalid key type.  Expecting a string")
		}

		if ks == "#Reason" {
			reason, ok = v.(int)

			if !ok {
				return 0, errors.New("unable to convert field value to integer")
			}

			continue
		}

		fkey := key.FKeyFor(ks)
		if fkey == key.INVALID_FIELDNAME {
			return 0, errors.New("invalid field name")
		}

		var found field.Field
		for _, f := range r.fields {
			if f.fkey == fkey {
				found = f.field
				break
			}
		}

		if found == nil {
			return 0, errors.New("no matching field name in primitive")
		}

		err := found.IngestValue(v)
		if err != nil {
			return 0, err
		}
	}

	return reason, nil
}
