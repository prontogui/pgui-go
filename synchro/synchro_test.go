package synchro

import (
	"reflect"
	"testing"

	"github.com/prontogui/golib/primitive"
	// "github.com/prontogui/golib/testhelp"
)

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

func verifyListOfPrimitives(t *testing.T, primitives []any, expecting ...primitive.Primitive) {

	len_p, len_e := len(primitives), len(expecting)

	if len_p != len_e {
		t.Fatalf("there are %d primitives.  Expecting %d.", len_p, len_e)
	}

	for i := 0; i < len(primitives); i++ {
		if !reflect.DeepEqual(primitives[i], expecting[i]) {
			t.Fatalf("primitive element %d is not equal to what's expected", i)
		}
	}
}

func Test_FullUpdate(t *testing.T) {

	top := []any{&primitive.Command{}}
	s := NewSynchro()
	s.SetTopPrimitives(top)
	// Verify there is a full update pending
	update := s.GetFullUpdate()
	primitives := verifyUpdateStructure(t, update, true)

	ec := &primitive.Command{}

	verifyListOfPrimitives(t, primitives, ec)
}

// TODO
// Build a function that returns a somewhat sophisticated list of primitives for testing partial updates.
// Write tests to see if partial updates are created currectly as individual and set of fields are changed.
//
