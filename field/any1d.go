package field

import (
	"github.com/prontogui/golib/primitive"
)

type Any1D struct {
	Reserved
	aa []primitive.Interface
}

func (f *Any1D) Get() []primitive.Interface {
	return f.aa
}

func (f *Any1D) Set(aa []primitive.Interface) {
	f.aa = aa
	// TODO:  prepare all the primitives for updates
	f.OnSet(true)
}
