package field

import (
	"errors"

	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

type Any struct {
	Reserved
	p primitive.Interface
}

func (f *Any) Get() primitive.Interface {
	return f.p
}

func (f *Any) Set(p primitive.Interface) {
	f.p = p
	f.OnSet(true)
}

func (f *Any) PrepareForUpdates(fkey key.FKey, pkey key.PKey, onset key.OnSetFunction, nextContainerIndex int) (isContainer bool) {

	isContainer = true

	f.StashUpdateInfo(fkey, pkey, onset)
	if f.p != nil {
		f.p.PrepareForUpdates(pkey.AddLevel(nextContainerIndex), onset)
	}

	return
}

func (f *Any) EgestValue() any {
	if f.p != nil {
		return f.p.EgestUpdate(true, nil)
	} else {
		return nil
	}
}

func (f *Any) IngestValue(value any) error {

	m, ok := value.(map[any]any)
	if !ok {
		return errors.New("invalid update")
	}

	if f.p != nil {
		return f.p.IngestUpdate(m)
	}

	return nil
}
