package synchro

import (
	"reflect"
	"testing"

	cbor "github.com/fxamacker/cbor/v2"
	"github.com/prontogui/golib/golib"
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
	s.SetTopPrimitives(&golib.Command{})

	// Verify there is a full update pending
	ec := &golib.Command{}
	verifyFullUpdate(t, s.GetFullUpdate(), ec)
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

	cmd1 := &golib.Command{}
	cmd2 := &golib.Command{}
	cmd3 := &golib.Command{}

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
	if len != 4 {
		t.Fatalf("partial update returned %d items.  Expecting 4 items", len)
	}

	verifyUpdateItemPKey(t, updates[0], 0)

	m1 := map[string]any{"Label": "Guten Tag!", "Issued": true}
	verifyUpdateItemMap(t, updates[1], m1)

	verifyUpdateItemPKey(t, updates[2], 2)

	m2 := map[string]any{"Status": uint64(2)}
	verifyUpdateItemMap(t, updates[3], m2)
}

// TODO
// Build a function that returns a somewhat sophisticated list of primitives for testing partial updates.
// Write tests to see if partial updates are created currectly as individual and set of fields are changed.
//
