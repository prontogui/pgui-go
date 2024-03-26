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

func (f *Any1D) Get() []primitive.Interface {
	return f.ary
}

func (f *Any1D) Set(ary []primitive.Interface) {
	f.ary = ary
	f.OnSet(true)
}

func (f *Any1D) PrepareForUpdates(fkey key.FKey, pkey key.PKey, onset key.OnSetFunction) {

	f.StashUpdateInfo(fkey, pkey, onset)

	// Prepare the children too
	for i, p := range f.ary {
		p.PrepareForUpdates(pkey.AddLevel(i), onset)
	}
}

func (f *Any1D) EgestValue() any {

	ary := []any{}

	for _, v := range f.ary {
		ary = append(ary, v.EgestUpdate(true, nil))
	}

	return ary
}

func (f *Any1D) IngestValue(value any) error {
	return errors.New("ingesting value for Any1D is not supported")
}
