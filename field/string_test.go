package field

import (
	"testing"

	"github.com/prontogui/golib/key"
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

	f.PrepareForUpdates(10, key.NewPKey(50), getTestOnsetFunc(), 0)

	verifyStashUpdateInfo(t, &f.Reserved)

	f.Set("abc")

	if !testOnsetCalled {
		t.Error("onset was not called")
	}
}

func Test_StringEgestValue(t *testing.T) {
	f := String{}
	f.Set("yabadabadoo!")
	v := f.EgestValue()
	s, ok := v.(string)
	if !ok {
		t.Fatal("cannot convert return value to string")
	}
	if s != "yabadabadoo!" {
		t.Fatal("incorrect value returned")
	}
}

func Test_StringIngestUpdate(t *testing.T) {

	f := String{}
	f.PrepareForUpdates(10, key.NewPKey(50), getTestOnsetFunc(), 0)

	err := f.IngestValue("hello, darling")

	testfunc := func() bool {
		return f.Get() == "hello, darling"
	}

	verifyIngestUpdateSuccessful(t, err, testfunc)
}

func Test_StringIngestUpdateEmptyString(t *testing.T) {

	f := String{}
	f.Set("goodbye, dear")
	f.PrepareForUpdates(10, key.NewPKey(50), getTestOnsetFunc(), 0)

	err := f.IngestValue("")

	testfunc := func() bool {
		return f.Get() == ""
	}

	verifyIngestUpdateSuccessful(t, err, testfunc)
}

func Test_StringIngestUpdateInvalid(t *testing.T) {

	f := String{}
	err := f.IngestValue(false)
	verifyIngestUpdateInvalid(t, err)
}
