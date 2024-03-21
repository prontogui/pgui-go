package field

import (
	"reflect"
	"testing"
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

	f.PrepareForUpdates("Abc", 50, testOnset)

	verifyStashUpdateInfo(t, &f.Reserved)

	f.Set([]byte{1, 2, 3})

	if !testOnsetCalled {
		t.Error("onset was not called")
	}
}
