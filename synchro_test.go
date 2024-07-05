// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"reflect"
	"testing"

	cbor "github.com/fxamacker/cbor/v2"
	"github.com/prontogui/golib/key"
)

func verifyFullUpdate(t *testing.T, cborUpdate []byte, expecting ...*SimplePrimitive) {

	if cborUpdate == nil {
		t.Fatal("no update (nil) was returned.  Expecting a CBOR-encoded update.")
	}

	var update any
	err := cbor.Unmarshal(cborUpdate, &update)

	if err != nil {
		t.Fatalf("attempt to unmarshall the CBOR encoded update resulted in error:  %s", err.Error())
	}

	updateList, ok := update.([]any)
	if !ok {
		t.Fatal("the returned update has invalid structure.")
	}
	if len(updateList) < 1 {
		t.Fatal("the update returned a list with wrong number of items.  Expecting at least 1 items.")
	}

	flag, ok := updateList[0].(bool)

	if !ok {
		t.Fatal("first elemenent of returned udpate is not a boolean.")
	}

	if !flag {
		t.Fatal("partial update returned.  Expecting a full update to be returned.")
	}

	len_p, len_e := len(updateList)-1, len(expecting)

	if len_p != len_e {
		t.Fatalf("there are %d items in update.  Expecting %d.", len_p, len_e)
	}

	for i, ulitem := range updateList[1:] {

		// Narrow down to a map
		m1 := ulitem.(map[any]any)
		tc := expecting[i]

		// Does every field in testcommand equal the same value in map? (this should be commutative)
		if tc.Label.Get() != m1["Label"].(string) {
			t.Fatalf("update item %d is not equal to what's expected", i)
		}
		if uint64(tc.Status.Get()) != m1["Status"].(uint64) {
			t.Fatalf("update item %d is not equal to what's expected", i)
		}
		if tc.Issued.Get() != m1["Issued"].(bool) {
			t.Fatalf("update item %d is not equal to what's expected", i)
		}

	}

}

func Test_FullUpdate(t *testing.T) {

	s := NewSynchro()
	s.SetTopPrimitives(&SimplePrimitive{})

	// Verify there is a full update pending
	ec := &SimplePrimitive{}
	fullupdate, err := s.GetFullUpdate()
	if err != nil {
		t.Fatalf("unexpected error:  %s", err.Error())
	}
	verifyFullUpdate(t, fullupdate, ec)
}

func verifyUpdateItemFalse(t *testing.T, item any) {
	flag, ok := item.(bool)
	if !ok {
		t.Fatal("update flag cannot be converted to bool")
	}

	if flag == true {
		t.Fatal("update flag is true.  Expecting a flag of false to indicate partial update")
	}
}

func verifyUpdateItemPKey(t *testing.T, item any, pkey key.PKey) {
	itemPKeyAny, ok := item.([]any)
	if !ok {
		t.Fatal("update item cannot be converted to PKey")
	}

	itemPKey := key.NewPKeyFromAny(itemPKeyAny...)
	if !key.NewPKey(itemPKey...).EqualTo(pkey) {
		t.Fatal("update item does not match expected PKey")
	}
}

func verifyUpdateItemMap(t *testing.T, item any, m map[string]any) {
	itemmap, ok := item.(map[any]any)
	if !ok {
		t.Fatal("update item is not map[any]any type")
	}

	if len(itemmap) != len(m) {
		t.Fatal("update item map is different size than expected")
	}

	for k, v := range m {
		v2, ok := itemmap[k]
		if !ok {
			t.Fatalf("key %s not found in update item map", k)
		}

		if !reflect.DeepEqual(v, v2) {
			t.Fatalf("update item key/value pair for '%s' does not match as expected", k)
		}
	}

}

func Test_PartialUpdate1(t *testing.T) {

	cmd1 := &SimplePrimitive{}
	cmd2 := &SimplePrimitive{}
	cmd3 := &SimplePrimitive{}

	s := NewSynchro()
	s.SetTopPrimitives(cmd1, cmd2, cmd3)

	// Test for no partial update yet
	pu, err := s.GetPartialUpdate()
	if pu != nil {
		t.Fatal("partial update available when nothing changed. Not expecting a partial update.")
	}
	if err != nil {
		t.Fatalf("unexpected error %s while getting partial update", err.Error())
	}

	// Change command label
	cmd1.Label.Set("Guten Tag!")
	cmd1.Issued.Set(true)

	cmd3.Status.Set(2)

	// Test for partial updates available
	updatesCbor, err := s.GetPartialUpdate()
	if updatesCbor == nil {
		t.Fatal("no updates available")
	}
	if err != nil {
		t.Fatalf("unexpected error %s while getting partial update", err.Error())
	}

	// Verify the content of updates
	var updates []any
	err = cbor.Unmarshal(updatesCbor, &updates)
	if err != nil {
		t.Fatalf("attempt to unmarshall updateds resulted in error of %s", err.Error())
	}

	len := len(updates)
	if len != 5 {
		t.Fatalf("partial update returned %d items.  Expecting 5 items", len)
	}

	verifyUpdateItemFalse(t, updates[0])

	verifyUpdateItemPKey(t, updates[1], key.NewPKey(0))

	m1 := map[string]any{"Label": "Guten Tag!", "Issued": true}
	verifyUpdateItemMap(t, updates[2], m1)

	verifyUpdateItemPKey(t, updates[3], key.NewPKey(2))

	m2 := map[string]any{"Status": uint64(2)}
	verifyUpdateItemMap(t, updates[4], m2)
}

func verifyPrimitivesEqual(t *testing.T, a []Primitive, b []Primitive) {

	lena, lenb := len(a), len(b)

	if lena != lenb {
		t.Fatalf("first set of primitives (a) has length of %d and second set (b) has length of %d.  Expecting equal number of primitives", lena, lenb)
	}

	for i, p := range a {

		sp1 := p.(*SimplePrimitive)
		sp2 := b[i].(*SimplePrimitive)

		if sp1.Label.Get() != sp2.Label.Get() {
			t.Errorf("Label fields of primitives a[%d] and b[%d] are not equal", i, i)
		}
		if sp1.Status.Get() != sp2.Status.Get() {
			t.Errorf("Status fields of primitives a[%d] and b[%d] are not equal", i, i)
		}
		if sp1.Issued.Get() != sp2.Issued.Get() {
			t.Errorf("Issued fields of primitives a[%d] and b[%d] are not equal", i, i)
		}
	}
}

func Test_IngestFullUpdateNotSupported(t *testing.T) {

	cmd1 := &SimplePrimitive{}
	cmd2 := &SimplePrimitive{}
	cmd3 := &SimplePrimitive{}

	s1 := NewSynchro()
	s1.SetTopPrimitives(cmd1, cmd2, cmd3)

	fullupdate, _ := s1.GetFullUpdate()

	s2 := NewSynchro()
	_, err := s2.IngestUpdate(fullupdate)

	if err == nil {
		t.Fatal("ingestion of full update should not be supported at this time")
	}

	errMsg := err.Error()
	expMsg := "ingestion of full updates is not supported"
	if errMsg != expMsg {
		t.Fatalf("a different error than expected was returned:  '%s'.  Expecting:  '%s'", errMsg, expMsg)
	}
}

func Test_IngestPartialUpdateMultiplePrimitivesNotSupported(t *testing.T) {
	cmd1 := &SimplePrimitive{}
	cmd2 := &SimplePrimitive{}
	cmd3 := &SimplePrimitive{}

	s1 := NewSynchro()
	s1.SetTopPrimitives(cmd1, cmd2, cmd3)

	cmd1.Issued.Set(true)
	cmd2.Label.Set("blah blah")

	var err error

	partialupdate, err := s1.GetPartialUpdate()
	if err != nil {
		t.Fatal("error getting a partial update from synchro")
	}

	s2 := NewSynchro()
	s2.SetTopPrimitives(&SimplePrimitive{}, &SimplePrimitive{}, &SimplePrimitive{})

	_, err = s2.IngestUpdate(partialupdate)
	if err == nil {
		t.Fatalf("expecting an error when attempting a partial update of mulitple primitives")
	}

	if err.Error() != "partial update is limited to one primitive" {
		t.Fatalf("wrong error was returned")
	}
}

func Test_IngestPartialUpdate(t *testing.T) {
	cmd1 := &SimplePrimitive{}
	cmd2 := &SimplePrimitive{}

	s1 := NewSynchro()
	s1.SetTopPrimitives(cmd1, cmd2)

	cmd2.Label.Set("blah blah")

	var err error

	partialupdate, err := s1.GetPartialUpdate()
	if err != nil {
		t.Fatal("error getting a partial update from synchro")
	}

	s2 := NewSynchro()
	s2.SetTopPrimitives(&SimplePrimitive{}, &SimplePrimitive{})

	_, err = s2.IngestUpdate(partialupdate)
	if err != nil {
		t.Fatalf("IngestUpdate returned error:  %s", err.Error())
	}

	// Are top primitives the same?
	verifyPrimitivesEqual(t, s1.GetTopPrimitives(), s2.GetTopPrimitives())
}
