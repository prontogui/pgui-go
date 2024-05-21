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

func (f *String) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset)
}

func (f *String) EgestValue() any {
	return f.s
}

func (f *String) IngestValue(value any) error {

	s, ok := value.(string)

	if !ok {
		return errors.New("unable to convert value (any) to field value")
	}

	f.s = s

	return nil
}
