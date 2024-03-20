package synchro

import (
	"reflect"
	"testing"

	cbor "github.com/fxamacker/cbor/v2"
	"github.com/prontogui/golib/coreprimitives"
	// "github.com/prontogui/golib/testhelp"
)

/* Version 1 - delete if all goes well.

func verifyUpdateStructure(t *testing.T, update any, isfull bool) []any {

	if update == nil {
		t.Fatal("no update returned.")
	}

	l, ok := update.([]any)
	if !ok {
		t.Fatal("the returned update has invalid structure.")
	}
	if len(l) != 2 {
		t.Fatal("the update returned a list with wrong number of items.  Expecting 2 items.")
	}

	flag, ok := l[0].(bool)

	if !ok {
		t.Fatal("first elemenent of returned udpate is not a boolean.")
	}
	if isfull {
		if !flag {
			t.Fatal("partial update returned.  Expecting a full update to be returned.")
		}
	} else {
		if flag {
			t.Fatal("full update returned.  Expecting a partial update to be returned.")
		}
	}

	updateItem, ok := l[1].([]any)

	if !ok {
		t.Fatal("second element of returned update is not a list of interfaces.")
	}

	return updateItem
}
*/

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
	s.SetTopPrimitives(&coreprimitives.Command{})

	// Verify there is a full update pending
	ec := &coreprimitives.Command{}
	verifyFullUpdate(t, s.GetFullUpdate(), ec)
}

func Test_PartialUpdate(t *testing.T) {

	cmd := &coreprimitives.Command{}
	cmd.Label.Set("Hello, World!")

	s := NewSynchro()
	s.SetTopPrimitives(cmd)

	// Test for no partial update yet
	if s.GetPartialUpdate() != nil {
		t.Fatal("partial update available when nothing changed. Not expecting a partial update.")
	}

	// Change command label
	cmd.Label.Set("Guten Tag!")

	// Test for partial updates available
	updatesCbor := s.GetPartialUpdate()
	if updatesCbor == nil {
		t.Fatal("no updates available")
	}

	// Verify the content of updates
	var updates []any
	err := cbor.Unmarshal(updatesCbor, &updates)
	if err != nil {
		t.Fatalf("attempt to unmarshall updateds resulted in error of %s", err.Error())
	}

}

// TODO
// Build a function that returns a somewhat sophisticated list of primitives for testing partial updates.
// Write tests to see if partial updates are created currectly as individual and set of fields are changed.
//
