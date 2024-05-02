package field

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_EventSetAndGetTrue(t *testing.T) {
	f := Event{}

	f.Set(true)

	if f.Get() != false {
		t.Fatal("cannot set boolean to true and get false value back.")
	}
}

func Test_EventPrepareForUpdates(t *testing.T) {
	f := Event{}

	f.PrepareForUpdates(10, key.NewPKey(50), getTestOnsetFunc())

	f.Set(true)

	if testOnsetCalled {
		t.Error("onset was called.  This is not expected.")
	}
}

func Test_EventEgestValue(t *testing.T) {
	f := Event{}
	f.Set(true)
	v := f.EgestValue()
	if v != nil {
		t.Fatal("egest value is not nil. Not expecting a value to be returned.")
	}
}

func Test_EventIngestUpdateTrue(t *testing.T) {

	f := Event{}

	err := f.IngestValue(true)

	testfunc := func() bool {
		return f.Get() == true
	}

	verifyIngestUpdateSuccessful(t, err, testfunc)
}

func Test_EventIngestUpdateFalse(t *testing.T) {

	f := Event{}
	f.Set(true)

	err := f.IngestValue(false)

	testfunc := func() bool {
		return f.Get() == false
	}

	verifyIngestUpdateSuccessful(t, err, testfunc)
}

func Test_EventIngestUpdateInvalid(t *testing.T) {

	f := Event{}
	err := f.IngestValue(10)
	verifyIngestUpdateInvalid(t, err)
}
