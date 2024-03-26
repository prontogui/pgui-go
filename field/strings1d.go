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

func (f *Strings1D) PrepareForUpdates(fkey key.FKey, pkey key.PKey, onset key.OnSetFunction) {
	f.StashUpdateInfo(fkey, pkey, onset)
}

func (f *Strings1D) EgestValue() any {
	return nil
}

func (f *Strings1D) IngestValue(value any) error {
	return errors.New("ingesting value for Strings1D is not supported")
}
