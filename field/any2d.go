package field

import (
	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

// TODO:  swap any type with primitive.Interface and update the test accordingly.
type Any2D struct {
	Reserved
	ary [][]primitive.Interface
}

func (f *Any2D) Get() [][]primitive.Interface {
	return f.ary
}

func (f *Any2D) Set(ary [][]primitive.Interface) {
	f.ary = ary
	f.OnSet(true)
}

func (f *Any2D) PrepareForUpdates(fieldname string, pkey key.PKey, onset key.OnSetFunction) {

	f.StashUpdateInfo(fieldname, pkey, onset)

	// Prepare the children too
	for i, p1 := range f.ary {
		pkeyi := pkey.AddLevel(i)

		for j, p2 := range p1 {
			pkeyj := pkeyi.AddLevel(j)
			p2.PrepareForUpdates(pkeyj, onset)
		}
	}
}
