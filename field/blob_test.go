// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package field

import (
	"reflect"
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_BlobSetAndGet(t *testing.T) {
	f := Blob{}

	f.Set([]byte{34, 200, 90, 1, 0})

	if !reflect.DeepEqual(f.Get(), []byte{34, 200, 90, 1, 0}) {
		t.Fatal("cannot set blob and get the same value back.")
	}
}

func Test_BlobPrepareForUpdates(t *testing.T) {
	f := Blob{}

	f.PrepareForUpdates(10, key.NewPKey(50), 0, getTestOnsetFunc())

	verifyFieldPreppedForUpdate(t, &f.Reserved)

	f.Set([]byte{1, 2, 3})

	if !testOnsetCalled {
		t.Error("onset was not called")
	}
}

func Test_BlobEgestValue(t *testing.T) {

	f := Blob{}
	f.Set([]byte{10, 20, 30})

	v := f.EgestValue()
	ba, ok := v.([]byte)
	if !ok {
		t.Fatal("unable to convert value to []byte")
	}
	if !reflect.DeepEqual(ba, []byte{10, 20, 30}) {
		t.Fatal("incorrect value returned")
	}
}

func Test_BlobIngestUpdate(t *testing.T) {

	f := Blob{}
	err := f.IngestValue([]byte{})
	if err == nil || err.Error() != "ingesting value for Blob is not supported" {
		t.Fatal("ingesting value for Blob should not be supported yet")
	}
}
