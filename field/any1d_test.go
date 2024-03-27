package field

import (
	"reflect"
	"testing"

	"github.com/prontogui/golib/primitive"
)

func Test_Any1DSetAndGet(t *testing.T) {
	f := Any1D{}

	actuals_i, _ := generateTestData1D()
	f.Set(actuals_i)

	expected_i, _ := generateTestData1D()

	if !reflect.DeepEqual(f.Get(), expected_i) {
		t.Fatal("cannot set any array and get the same value back.")
	}
}

func Test_Any1DPrepareForUpdates(t *testing.T) {
	f := Any1D{}

	values_i, values_p := generateTestData1D()

	f.Set(values_i)

	f.PrepareForUpdates(10, 50, getTestOnsetFunc())

	verifyStashUpdateInfo(t, &f.Reserved)

	for i, p := range values_p {
		if !p.prepped {
			t.Errorf("array element (%d) was not prepared correctly", i)
		}
	}

	f.Set(values_i)

	if !testOnsetCalled {
		t.Error("onset was not called")
	}
}

func Test_Any1DEgestValue(t *testing.T) {
	f := Any1D{}
	f.Set([]primitive.Interface{&TestPrimitive{s: "abc"}, &TestPrimitive{s: "xyz"}})
	v := f.EgestValue()
	a, ok := v.([]any)
	if !ok {
		t.Fatal("cannot convert value to []any")
	}
	if len(a) != 2 {
		t.Fatal("wrong number of elements returned.  Expecting 2 elements")
	}
	m1, ok := a[0].(map[string]any)
	if !ok {
		t.Fatal("cannot convert element to map[string]any")
	}
	m1v, ok := m1["s"].(string)
	if !ok {
		t.Fatal("cannot convert element map item to string")
	}
	if m1v != "abc" {
		t.Fatal("wrong string value for element")
	}
	m2, ok := a[1].(map[string]any)
	if !ok {
		t.Fatal("cannot convert element to map[string]any")
	}
	m2v, ok := m2["s"].(string)
	if !ok {
		t.Fatal("cannot convert element map item to string")
	}
	if m2v != "xyz" {
		t.Fatal("wrong string value for element")
	}
}

/*
Prelimary thoughts on how this should work...
* - it should be possible to add one or more primitives to an existing array, e.g. adding a Text item to a GUI.
* - new items must be added at the end of the list, after items that were created by server code.
* - it must be possible to infer the primitive type from the signature of fields in the update item.  I want to
*   avoid adding a type identifier field, since it really isn't necessary and its redundant information.
* - deleting a primitive, that was created by the app, may be possible.  Deleting primitives created by the server
*   cannot be possible, since it would break code interacting with it.   Maybe app created primitives could be
*   marked somehow, e.g. using a reserved field like "delete-me".  A partial update could be streamed to the app once
*   deletions are made on the server end.
*/
func Test_Any1DIngestUpdate(t *testing.T) {
	/*
	   f := Any1D{}
	   f.Set([]primitive.Interface{&TestPrimitive{s: "a"}, &TestPrimitive{s: "b"}})

	   err := f.IngestValue([]string{"abc", "def"})

	   	if err != nil {
	   		t.Fatalf("unexpected error was returned:  %s", err.Error())
	   	}

	   	if !reflect.DeepEqual(f.Get(), []string{"abc", "def"}) {
	   		t.Fatal("value not set correctly")
	   	}
	*/
}

func Test_Any1DIngestUpdateInvalid(t *testing.T) {
	/*
	   f := Strings1D{}
	   err := f.IngestValue(450)

	   	if err == nil {
	   		t.Fatal("error was not returned")
	   	}

	   	if err.Error() != "cannot convert value to []string" {
	   		t.Fatal("wrong error was returned")
	   	}
	*/
}
