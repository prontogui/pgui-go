package primitive

import "github.com/prontogui/golib/key"

type Interface interface {
	PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction)
	GetChildPrimitive(index int) Interface
	GetFieldValue(fkey key.FKey) any
	IngestFieldUpdate(fieldname string, update any) error
}
