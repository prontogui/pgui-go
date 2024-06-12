// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package field

import (
	"reflect"
	"testing"

	"github.com/prontogui/golib/key"
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

	f.PrepareForUpdates(10, key.NewPKey(50), 0, getTestOnsetFunc())

	verifyFieldPreppedForUpdate(t, &f.Reserved)

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
	err := f.IngestValue([]string{"abc", "def"})
	if err != nil {
		t.Fatalf("unexpected error was returned:  %s", err.Error())
	}
	if !reflect.DeepEqual(f.Get(), []string{"abc", "def"}) {
		t.Fatal("value not set correctly")
	}
}

func Test_Strings1DIngestUpdateInvalid(t *testing.T) {

	f := Strings1D{}
	err := f.IngestValue(450)
	if err == nil {
		t.Fatal("error was not returned")
	}
	if err.Error() != "cannot convert value to []string" {
		t.Fatal("wrong error was returned")
	}
}
