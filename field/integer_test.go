package field

import (
	"testing"

	"github.com/prontogui/golib/key"
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

	f.PrepareForUpdates(10, key.NewPKey(50), getTestOnsetFunc())

	verifyStashUpdateInfo(t, &f.Reserved)

	f.Set(92342)

	if !testOnsetCalled {
		t.Error("onset was not called")
	}
}

func Test_IntegerEgestValue(t *testing.T) {
	f := Integer{}
	f.Set(12345)
	v := f.EgestValue()
	i, ok := v.(int)
	if !ok {
		t.Fatal("cannot convert return value to int")
	}
	if i != 12345 {
		t.Fatal("incorrect value returned")
	}
}

func Test_IntegerIngestUpdate(t *testing.T) {

	f := Integer{}
	f.PrepareForUpdates(10, key.NewPKey(50), getTestOnsetFunc())

	err := f.IngestValue(3400)

	testfunc := func() bool {
		return f.Get() == 3400
	}

	verifyIngestUpdateSuccessful(t, err, testfunc)
}

func Test_IntegerIngestUpdateZero(t *testing.T) {

	f := Integer{}
	f.Set(290)
	f.PrepareForUpdates(10, key.NewPKey(50), getTestOnsetFunc())

	err := f.IngestValue(0)

	testfunc := func() bool {
		return f.Get() == 0
	}

	verifyIngestUpdateSuccessful(t, err, testfunc)
}

func Test_IntegerIngestUpdateInvalid(t *testing.T) {

	f := Integer{}
	err := f.IngestValue(false)
	verifyIngestUpdateInvalid(t, err)
}
