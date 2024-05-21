package field

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type Strings1D struct {
	Reserved
	sa []string
}

func (f *Strings1D) Get() []string {
	return f.sa
}

func (f *Strings1D) Set(sa []string) {
	f.sa = sa
	f.OnSet(false)
}

func (f *Strings1D) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset)
}

func (f *Strings1D) EgestValue() any {
	return f.sa
}

func (f *Strings1D) IngestValue(value any) error {
	sa, ok := value.([]string)
	if !ok {
		return errors.New("cannot convert value to []string")
	}
	f.sa = sa
	return nil
}
