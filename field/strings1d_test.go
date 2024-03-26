package field

import (
	"reflect"
	"testing"
)

func Test_String1DSetAndGet(t *testing.T) {
	f := Strings1D{}

	sa := []string{"abc", "xyz", "def"}

	f.Set(sa)

	if !reflect.DeepEqual(f.Get(), sa) {
		t.Fatal("cannot set string array and get the same value back.")
	}
}

func Test_String1DPrepareForUpdates(t *testing.T) {
	f := Strings1D{}

	f.PrepareForUpdates(10, 50, getTestOnsetFunc())

	verifyStashUpdateInfo(t, &f.Reserved)

	f.Set([]string{"abc", "xyz"})

	if !testOnsetCalled {
		t.Error("onset was not called")
	}
}

func Test_Strings1DEgestValue(t *testing.T) {

	f := Strings1D{}
	f.Set([]string{"x", "y", "z"})

	v := f.EgestValue()
	sa, ok := v.([]string)
	if !ok {
		t.Fatal("unable to convert value to []string")
	}
	if !reflect.DeepEqual(sa, []string{"x", "y", "z"}) {
		t.Fatal("incorrect value returned")
	}
}

func Test_Strings1DIngestUpdate(t *testing.T) {

	f := Strings1D{}
	err := f.IngestValue([]byte{})
	if err == nil || err.Error() != "ingesting value for Strings1D is not supported" {
		t.Fatal("ingesting value for Strings1D should not be supported yet")
	}
}
