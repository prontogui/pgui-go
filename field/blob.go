package field

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type Blob struct {
	Reserved
	blob []byte
}

func (f *Blob) Get() []byte {
	return f.blob
}

func (f *Blob) Set(blob []byte) {
	f.blob = blob
	f.OnSet(false)
}

func (f *Blob) PrepareForUpdates(fieldname string, pkey key.PKey, onset key.OnSetFunction) {
	f.StashUpdateInfo(fieldname, pkey, onset)
}

func (f *Blob) IngestUpdate(update any) error {
	return errors.New("ingesting field update for Blob is not supported")
}
