package field

import "github.com/prontogui/golib/key"

type Field interface {
	PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) (isContainer bool)
	EgestValue() any
	IngestValue(value any) error
}
