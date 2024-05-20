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

func (f *Blob) PrepareForUpdates(fkey key.FKey, pkey key.PKey, onset key.OnSetFunction, nextContainerIndex int) (isContainer bool) {
	isContainer = false
	f.StashUpdateInfo(fkey, pkey, onset)
	return
}

func (f *Blob) EgestValue() any {
	return f.blob
}

func (f *Blob) IngestValue(value any) error {
	return errors.New("ingesting value for Blob is not supported")
}
