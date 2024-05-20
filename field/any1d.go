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

func (f *Any1D) PrepareForUpdates(fkey key.FKey, pkey key.PKey, onset key.OnSetFunction, nextContainerIndex int) (isContainer bool) {

	isContainer = true

	f.StashUpdateInfo(fkey, pkey, onset)

	// Prepare the children too
	for i, p := range f.ary {
		p.PrepareForUpdates(pkey.AddLevel(nextContainerIndex).AddLevel(i), onset)
	}

	return
}

func (f *Any1D) EgestValue() any {

	ary := []any{}

	for _, v := range f.ary {
		ary = append(ary, v.EgestUpdate(true, nil))
	}

	return ary
}

func (f *Any1D) IngestValue(value any) error {

	l, ok := value.([]any)
	if !ok {
		return errors.New("invalid update")
	}

	if len(l) != len(f.ary) {
		return errors.New("number of primitives in update does not equal existing primitives")
	}

	for i, v := range l {
		m, ok := v.(map[any]any)
		if !ok {
			return errors.New("invalid update")
		}

		err := f.ary[i].IngestUpdate(m)
		if err != nil {
			return err
		}
	}

	return nil
}
