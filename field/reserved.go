package field

import (
	"github.com/prontogui/golib/key"
)

type Reserved struct {
	pkey  key.PKey // `cbor:"omitempty"`
	fkey  key.FKey
	onset func(key.PKey, key.FKey)
}

func (f *Reserved) PrepareForUpdates(fieldname string, pkey key.PKey, onset func(key.PKey, key.FKey)) {
	f.fkey = key.FKeyFor(fieldname)
	f.pkey = pkey
	f.onset = onset
}

func (f *Reserved) OnSet() {
	if f.onset != nil {
		f.onset(f.pkey, f.fkey)
	}
}
