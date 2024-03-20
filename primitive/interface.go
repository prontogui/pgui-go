package primitive

import "github.com/prontogui/golib/key"

type Interface interface {
	PrepareForUpdates(key.PKey, key.OnSetFunction)
}
