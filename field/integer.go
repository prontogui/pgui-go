package field

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type Integer struct {
	Reserved
	i int
}

func (f *Integer) Get() int {
	return f.i
}

func (f *Integer) Set(i int) {
	f.i = i
	f.OnSet(false)
}

func (f *Integer) PrepareForUpdates(fkey key.FKey, pkey key.PKey, onset key.OnSetFunction) {
	f.StashUpdateInfo(fkey, pkey, onset)
}

func (f *Integer) EgestValue() any {
	return f.i
}

func (f *Integer) IngestValue(value any) error {

	i, ok := value.(int)

	if !ok {
		return errors.New("unable to convert value (any) to field value")
	}

	f.i = i

	return nil
}
