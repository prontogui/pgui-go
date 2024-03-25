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

	f.PrepareForUpdates(10, 50, getTestOnsetFunc())

	verifyStashUpdateInfo(t, &f.Reserved)

	f.Set("abc")

	if !testOnsetCalled {
		t.Error("onset was not called")
	}
}

func Test_StringIngestUpdate(t *testing.T) {

	f := String{}
	f.PrepareForUpdates(10, 50, getTestOnsetFunc())

	err := f.IngestUpdate("hello, darling")

	testfunc := func() bool {
		return f.Get() == "hello, darling"
	}

	verifyIngestUpdateSuccessful(t, err, testfunc)
}

func Test_StringIngestUpdateEmptyString(t *testing.T) {

	f := String{}
	f.Set("goodbye, dear")
	f.PrepareForUpdates(10, 50, getTestOnsetFunc())

	err := f.IngestUpdate("")

	testfunc := func() bool {
		return f.Get() == ""
	}

	verifyIngestUpdateSuccessful(t, err, testfunc)
}

func Test_StringIngestUpdateInvalid(t *testing.T) {

	f := String{}
	err := f.IngestUpdate(false)
	verifyIngestUpdateInvalid(t, err)
}
