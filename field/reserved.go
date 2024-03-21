package field

import (
	"github.com/prontogui/golib/key"
)

type Reserved struct {
	pkey  key.PKey // `cbor:"omitempty"`
	fkey  key.FKey
	onset func(key.PKey, key.FKey, bool)
}

func (f *Reserved) StashUpdateInfo(fieldname string, pkey key.PKey, onset key.OnSetFunction) {
	f.fkey = key.FKeyFor(fieldname)
	f.pkey = pkey
	f.onset = onset
}

func (f *Reserved) OnSet(structural bool) {
	if f.onset != nil {
		f.onset(f.pkey, f.fkey, structural)
	}
}
