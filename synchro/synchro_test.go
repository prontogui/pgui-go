package synchro

import (
	"reflect"
	"testing"

	cbor "github.com/fxamacker/cbor/v2"
	"github.com/prontogui/golib/primitive"
	// "github.com/prontogui/golib/testhelp"
)

func verifyFullUpdate(t *testing.T, cborUpdate []byte, expecting ...any) {

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

	for i, v := range updateList[1:] {
		// Marshal both items to CBOR in order to compare them
		exp_c, _ := cbor.Marshal(expecting[i])
		actual_c, _ := cbor.Marshal(v)

		if !reflect.DeepEqual(actual_c, exp_c) {
			t.Fatalf("update item %d is not equal to what's expected", i)
		}
	}

}

func Test_FullUpdate(t *testing.T) {

	s := NewSynchro()
	s.SetTopPrimitives(&testcommand{})

	// Verify there is a full update pending
	ec := &testcommand{}
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

func verifyUpdateItemPKey(t *testing.T, item any, pkey uint64) {
	itemPKey, ok := item.(uint64)
	if !ok {
		t.Fatal("update item cannot be converted to PKey")
	}

	if itemPKey != pkey {
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

	cmd1 := &testcommand{}
	cmd2 := &testcommand{}
	cmd3 := &testcommand{}

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

	verifyUpdateItemPKey(t, updates[1], 0)

	m1 := map[string]any{"Label": "Guten Tag!", "Issued": true}
	verifyUpdateItemMap(t, updates[2], m1)

	verifyUpdateItemPKey(t, updates[3], 2)

	m2 := map[string]any{"Status": uint64(2)}
	verifyUpdateItemMap(t, updates[4], m2)
}

func verifyPrimitivesEqual(t *testing.T, a []primitive.Interface, b []primitive.Interface) {

	lena, lenb := len(a), len(b)

	if lena != lenb {
		t.Fatalf("first set of primitives (a) has length of %d and second set (b) has length of %d.  Expecting equal number of primitives", lena, lenb)
	}

	for i, p := range a {
		if !reflect.DeepEqual(p, b[i]) {
			t.Errorf("primitives a[%d] and b[%d] are not equal.  Expecting them to be identical", i, i)
		}
	}
}

func Test_IngestPartialUpdate(t *testing.T) {
	cmd1 := &testcommand{}
	cmd2 := &testcommand{}
	cmd3 := &testcommand{}

	s1 := NewSynchro()
	s1.SetTopPrimitives(cmd1, cmd2, cmd3)

	var err error

	fullupdate, err := s1.GetFullUpdate()
	if err != nil {
		t.Fatal("error getting a full update from synchro")
	}

	s2 := NewSynchro()
	err = s2.IngestPartialUpdates(fullupdate)
	if err != nil {
		t.Fatalf("IngestUpdates returned error:  %s", err.Error())
	}

	// Are top primitives the same?
	verifyPrimitivesEqual(t, s1.GetTopPrimitives(), s2.GetTopPrimitives())
}

// TODO
// Build a function that returns a somewhat sophisticated list of primitives for testing partial updates.
// Write tests to see if partial updates are created currectly as individual and set of fields are changed.
//
