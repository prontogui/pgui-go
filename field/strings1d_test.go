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

	f.PrepareForUpdates("Abc", 50, getTestOnsetFunc())

	verifyStashUpdateInfo(t, &f.Reserved)

	f.Set([]string{"abc", "xyz"})

	if !testOnsetCalled {
		t.Error("onset was not called")
	}
}

func Test_Strings1DIngestUpdate(t *testing.T) {

	f := Strings1D{}
	err := f.IngestUpdate([]byte{})
	if err == nil || err.Error() != "ingesting field update for Strings1D is not supported" {
		t.Fatal("ingesting update for Strings1D should not be supported yet")
	}
}
