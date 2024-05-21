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

func (f *Any) prepareDescendantForUpdate() {
	if f.p != nil {
		if f.onset == nil {
			f.p.PrepareForUpdates(key.EmptyPKey(), nil)
		} else {
			f.p.PrepareForUpdates(f.pkey.AddLevel(f.fieldPKeyIndex), f.onset)
		}
	}
}

func (f *Any) Get() primitive.Interface {
	return f.p
}

func (f *Any) Set(p primitive.Interface) {
	f.p = p
	f.prepareDescendantForUpdate()
	f.OnSet(true)
}

func (f *Any) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) (isContainer bool) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset)
	f.prepareDescendantForUpdate()
	return true
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
