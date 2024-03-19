package primitive

import "github.com/prontogui/golib/key"

const (
	// The maximum number of fields in any given primitive.  TODO:  check for accuracy of this in unit testing,
	// in case a primitive is updated or added without changing this number.
	MaxPrimitiveFields = 4
)

type Primitive interface {
	NotifyOnSet(key.PKey, func(key.PKey, key.FKey))
}
