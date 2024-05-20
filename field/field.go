package field

import "github.com/prontogui/golib/key"

type Field interface {
	PrepareForUpdates(fkey key.FKey, pkey key.PKey, onset key.OnSetFunction, nextContainerIndex int) (isContainer bool)
	EgestValue() any
	IngestValue(value any) error
}
