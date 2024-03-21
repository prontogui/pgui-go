package field

import (
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

func (f *Any1D) PrepareForUpdates(fieldname string, pkey key.PKey, onset key.OnSetFunction) {

	f.StashUpdateInfo(fieldname, pkey, onset)

	// Prepare the children too
	for i, p := range f.ary {
		p.PrepareForUpdates(pkey.AddLevel(i), onset)
	}
}
