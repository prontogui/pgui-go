package field

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type Boolean struct {
	Reserved
	b bool
}

func (f *Boolean) Get() bool {
	return f.b
}

func (f *Boolean) Set(b bool) {
	f.b = b
	f.OnSet(false)
}

func (f *Boolean) PrepareForUpdates(fieldname string, pkey key.PKey, onset key.OnSetFunction) {
	f.StashUpdateInfo(fieldname, pkey, onset)
}

func (f *Boolean) IngestUpdate(update any) error {

	b, ok := update.(bool)

	if !ok {
		return errors.New("unable to convert update (any) to field value")
	}

	f.b = b
	return nil
}
