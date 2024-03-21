package field

import (
	"github.com/prontogui/golib/key"
)

type Boolean struct {
	Reserved
	b bool
}

func (f *Boolean) Get() bool {
	return f.b
}

func (f *Boolean) Set(b bool) {
	f.b = b
	f.OnSet(false)
}

func (f *Boolean) PrepareForUpdates(fieldname string, pkey key.PKey, onset key.OnSetFunction) {
	f.StashUpdateInfo(fieldname, pkey, onset)
}
