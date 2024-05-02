package field

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type Event struct {
	Reserved
	b bool
}

func (f *Event) GetAsAny() any {
	return f.b
}

func (f *Event) Get() bool {
	return f.b
}

func (f *Event) Set(b bool) {
}

func (f *Event) PrepareForUpdates(fkey key.FKey, pkey key.PKey, onset key.OnSetFunction) {
}

func (f *Event) EgestValue() any {
	// Inform caller that this field does not egest a value
	return nil
}

func (f *Event) IngestValue(value any) error {

	b, ok := value.(bool)

	if !ok {
		return errors.New("unable to convert value (any) to field value")
	}

	f.b = b
	return nil
}
