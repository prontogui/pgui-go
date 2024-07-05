// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_StringSetAndGet(t *testing.T) {
	f := StringField{}

	f.Set("abc")

	if f.Get() != "abc" {
		t.Fatal("cannot set string and get the same value back.")
	}
}

func Test_StringPrepareForUpdates(t *testing.T) {
	f := StringField{}

	f.PrepareForUpdates(10, key.NewPKey(50), 0, getTestOnsetFunc())

	verifyFieldPreppedForUpdate(t, &f.FieldBase)

	f.Set("abc")

	if !testOnsetCalled {
		t.Error("onset was not called")
	}
}

func Test_StringEgestValue(t *testing.T) {
	f := StringField{}
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

	f := StringField{}
	f.PrepareForUpdates(10, key.NewPKey(50), 0, getTestOnsetFunc())

	err := f.IngestValue("hello, darling")

	testfunc := func() bool {
		return f.Get() == "hello, darling"
	}

	verifyIngestUpdateSuccessful(t, err, testfunc)
}

func Test_StringIngestUpdateEmptyString(t *testing.T) {

	f := StringField{}
	f.Set("goodbye, dear")
	f.PrepareForUpdates(10, key.NewPKey(50), 0, getTestOnsetFunc())

	err := f.IngestValue("")

	testfunc := func() bool {
		return f.Get() == ""
	}

	verifyIngestUpdateSuccessful(t, err, testfunc)
}

func Test_StringIngestUpdateInvalid(t *testing.T) {

	f := StringField{}
	err := f.IngestValue(false)
	verifyIngestUpdateInvalid(t, err)
}
