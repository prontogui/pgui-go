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

func (f *Integer) PrepareForUpdates(fieldname string, pkey key.PKey, onset key.OnSetFunction) {
	f.StashUpdateInfo(fieldname, pkey, onset)
}

func (f *Integer) IngestUpdate(update any) error {

	i, ok := update.(int)

	if !ok {
		return errors.New("unable to convert update (any) to field value")
	}

	f.i = i

	return nil
}
