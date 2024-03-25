package field

import (
	"testing"
)

func Test_IntegerSetAndGet(t *testing.T) {
	f := Integer{}

	f.Set(92342)

	if f.Get() != 92342 {
		t.Fatal("cannot set integer and get the same value back.")
	}
}

func Test_IntegerPrepareForUpdates(t *testing.T) {
	f := Integer{}

	f.PrepareForUpdates(0, 50, getTestOnsetFunc())

	verifyStashUpdateInfo(t, &f.Reserved)

	f.Set(92342)

	if !testOnsetCalled {
		t.Error("onset was not called")
	}
}

func Test_IntegerIngestUpdate(t *testing.T) {

	f := Integer{}
	f.PrepareForUpdates(0, 50, getTestOnsetFunc())

	err := f.IngestUpdate(3400)

	testfunc := func() bool {
		return f.Get() == 3400
	}

	verifyIngestUpdateSuccessful(t, err, testfunc)
}

func Test_IntegerIngestUpdateZero(t *testing.T) {

	f := Integer{}
	f.Set(290)
	f.PrepareForUpdates(0, 50, getTestOnsetFunc())

	err := f.IngestUpdate(0)

	testfunc := func() bool {
		return f.Get() == 0
	}

	verifyIngestUpdateSuccessful(t, err, testfunc)
}

func Test_IntegerIngestUpdateInvalid(t *testing.T) {

	f := Integer{}
	err := f.IngestUpdate(false)
	verifyIngestUpdateInvalid(t, err)
}
