package primitive

import "github.com/prontogui/golib/key"

type Primitive interface {
	NotifyOnSet(key.PKey, func(key.PKey, key.FKey))
}
