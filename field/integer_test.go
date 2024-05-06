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

func Test_IntegerIngestUpdatePositive(t *testing.T) {

	f := Integer{}
	f.PrepareForUpdates(10, key.NewPKey(50), getTestOnsetFunc())

	testfunc := func() bool {
		return f.Get() == 34
	}

	verifyIngestUpdateSuccessful(t, f.IngestValue(uint8(34)), testfunc)
	verifyIngestUpdateSuccessful(t, f.IngestValue(int8(34)), testfunc)
	verifyIngestUpdateSuccessful(t, f.IngestValue(uint16(34)), testfunc)
	verifyIngestUpdateSuccessful(t, f.IngestValue(int16(34)), testfunc)
	verifyIngestUpdateSuccessful(t, f.IngestValue(uint32(34)), testfunc)
	verifyIngestUpdateSuccessful(t, f.IngestValue(int32(34)), testfunc)
	verifyIngestUpdateSuccessful(t, f.IngestValue(uint64(34)), testfunc)
	verifyIngestUpdateSuccessful(t, f.IngestValue(int64(34)), testfunc)
	verifyIngestUpdateSuccessful(t, f.IngestValue(uint(34)), testfunc)
	verifyIngestUpdateSuccessful(t, f.IngestValue(int(34)), testfunc)
}

func Test_IntegerIngestUpdateNegative(t *testing.T) {

	f := Integer{}
	f.PrepareForUpdates(10, key.NewPKey(50), getTestOnsetFunc())

	testfunc := func() bool {
		return f.Get() == -34
	}

	verifyIngestUpdateSuccessful(t, f.IngestValue(int8(-34)), testfunc)
	verifyIngestUpdateSuccessful(t, f.IngestValue(int16(-34)), testfunc)
	verifyIngestUpdateSuccessful(t, f.IngestValue(int32(-34)), testfunc)
	verifyIngestUpdateSuccessful(t, f.IngestValue(int64(-34)), testfunc)
	verifyIngestUpdateSuccessful(t, f.IngestValue(int(-34)), testfunc)
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
