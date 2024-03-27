package field

import (
	"errors"

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

func (f *Any2D) PrepareForUpdates(fkey key.FKey, pkey key.PKey, onset key.OnSetFunction) {

	f.StashUpdateInfo(fkey, pkey, onset)

	// Prepare the children too
	for i, p1 := range f.ary {
		pkeyi := pkey.AddLevel(i)

		for j, p2 := range p1 {
			pkeyj := pkeyi.AddLevel(j)
			p2.PrepareForUpdates(pkeyj, onset)
		}
	}
}

func (f *Any2D) EgestValue() any {
	ary := [][]any{}

	for _, row := range f.ary {
		ary2 := []any{}

		for _, cell := range row {
			ary2 = append(ary2, cell.EgestUpdate(true, nil))
		}

		ary = append(ary, ary2)
	}

	return ary
}

func (f *Any2D) IngestValue(value any) error {

	ary, ok := value.([][]any)
	if !ok {
		return errors.New("invalid update")
	}

	if len(ary) != len(f.ary) {
		return errors.New("number of primitives in update does not equal existing primitives")
	}

	for i, row := range ary {

		if len(row) != len(f.ary) {
			return errors.New("number of primitives in update does not equal existing primitives")
		}

		for j, v := range row {

			m, ok := v.(map[string]any)
			if !ok {
				return errors.New("invalid update")
			}

			err := f.ary[i][j].IngestUpdate(m)
			if err != nil {
				return err
			}
		}
	}

	return nil

}
