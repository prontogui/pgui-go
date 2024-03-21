package field

import (
	"testing"
)

func Test_StringSetAndGet(t *testing.T) {
	f := String{}

	f.Set("abc")

	if f.Get() != "abc" {
		t.Fatal("cannot set string and get the same value back.")
	}
}

func Test_StringPrepareForUpdates(t *testing.T) {
	f := String{}

	f.PrepareForUpdates("Abc", 50, testOnset)

	verifyStashUpdateInfo(t, &f.Reserved)

	f.Set("abc")

	if !testOnsetCalled {
		t.Error("onset was not called")
	}
}
