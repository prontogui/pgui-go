package field

import "github.com/prontogui/golib/key"

type Field interface {
	GetAsAny() any
	PrepareForUpdates(fkey key.FKey, pkey key.PKey, onset key.OnSetFunction)
	IngestUpdate(update any) error
}
