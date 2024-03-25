package field

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type String struct {
	Reserved
	s string
}

func (f *String) Get() string {
	return f.s
}

func (f *String) Set(s string) {
	f.s = s
	f.OnSet(false)
}

func (f *String) PrepareForUpdates(fieldname string, pkey key.PKey, onset key.OnSetFunction) {
	f.StashUpdateInfo(fieldname, pkey, onset)
}

func (f *String) IngestUpdate(update any) error {

	s, ok := update.(string)

	if !ok {
		return errors.New("unable to convert update (any) to field value")
	}

	f.s = s

	return nil
}
