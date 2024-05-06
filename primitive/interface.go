package primitive

import "github.com/prontogui/golib/key"

type Interface interface {
	PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction)
	GetChildPrimitive(index int) Interface
	EgestUpdate(fullupdate bool, fkeys []key.FKey) map[any]any
	IngestUpdate(update map[any]any) (int, error)
}
