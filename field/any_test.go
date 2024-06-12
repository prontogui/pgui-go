// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package field

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_AnySetAndGet(t *testing.T) {
	f := Any{}

	f.Set(&TestPrimitive{s: "abc"})

	tp, ok := f.Get().(*TestPrimitive)

	if !ok || tp.s != "abc" {
		t.Fatal("cannot set value and get the same value back.")
	}
}

func Test_AnySetWithFieldUnpreppedAndChildrenUnprepped(t *testing.T) {
	f := Any{}

	tp := &TestPrimitive{s: "abc"}
	f.Set(tp)

	verifyChildNotPreppedForUpdate(t, tp)
}

func Test_AnySetWithFieldUnpreppedAndChildrenPreviouslyPrepped(t *testing.T) {
	f := Any{}

	tp := &TestPrimitive{s: "abc"}

	bogeyPkey := key.NewPKey(1, 2, 3)
	bokeyOnset := getBogeyOnsetFunc()
	tp.PrepareForUpdates(bogeyPkey, bokeyOnset)

	f.Set(tp)

	verifyChildNotPreppedForUpdate(t, tp)
}

func Test_AnySetWithFieldPreppedAndChildrenUnprepped(t *testing.T) {
	f := Any{}

	pkey := key.NewPKey(50)
	onset := getTestOnsetFunc()
	f.PrepareForUpdates(10, pkey, 0, onset)

	tp := &TestPrimitive{s: "abc"}
	f.Set(tp)

	verifyChildPreppedForUpdate(t, tp, pkey.AddLevel(0), onset)
}

func Test_AnySetWithFieldPreppedAndChildrenPreviouslyPrepped(t *testing.T) {
	f := Any{}

	pkey := key.NewPKey(50)
	onset := getTestOnsetFunc()
	f.PrepareForUpdates(10, pkey, 0, onset)

	tp := &TestPrimitive{s: "abc"}

	bogeyPkey := key.NewPKey(1, 2, 3)
	bokeyOnset := getBogeyOnsetFunc()
	tp.PrepareForUpdates(bogeyPkey, bokeyOnset)

	f.Set(tp)

	verifyChildPreppedForUpdate(t, tp, pkey.AddLevel(0), onset)
}

func Test_AnyPrepareForUpdates(t *testing.T) {
	f := Any{}

	f.Set(&TestPrimitive{s: "abc"})

	f.PrepareForUpdates(10, key.NewPKey(50), 0, getTestOnsetFunc())

	verifyFieldPreppedForUpdate(t, &f.Reserved)

	f.Set(&TestPrimitive{s: "xyz"})

	if !testOnsetCalled {
		t.Error("onset was not called")
	}
}

func Test_AnyEgestValue(t *testing.T) {
	f := Any{}
	f.Set(&TestPrimitive{s: "abc"})
	v := f.EgestValue()
	_, ok := v.(map[any]any)
	if !ok {
		t.Fatal("cannot convert element to map[any]any")
	}
}

func Test_AnyIngestUpdate(t *testing.T) {

	f := &Any{}
	tp := &TestPrimitive{}
	f.Set(tp)

	m := map[any]any{"s": "Hello"}

	err := f.IngestValue(m)
	if err != nil {
		t.Fatalf("unexpected error returned:  %s", err.Error())
	}

	if tp.s != "Hello" {
		t.Fatal("primitive #1 not updated correctly")
	}
}

func Test_AnyIngestUpdateInvalid1(t *testing.T) {

	f := Any{}
	f.Set(&TestPrimitive{})

	err := f.IngestValue(3453)
	if err == nil {
		t.Fatal("no error returned for an invalid update")
	}
	if err.Error() != "invalid update" {
		t.Fatalf("wrong error was returned:  %s", err.Error())
	}
}

func Test_AnyIngestUpdateInvalid2(t *testing.T) {

	f := Any{}
	f.Set(&TestPrimitive{})

	err := f.IngestValue([]any{"Hello", "World"})

	if err == nil {
		t.Fatal("no error returned for an invalid update")
	}
	if err.Error() != "invalid update" {
		t.Fatalf("wrong error was returned:  %s", err.Error())
	}
}
