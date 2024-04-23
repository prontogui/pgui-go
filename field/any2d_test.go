package field

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

func Test_Any2DSetAndGet(t *testing.T) {
	f := Any2D{}

	actual, _ := generateTestData2D()
	f.Set(actual)

	expected, _ := generateTestData2D()

	if !reflect.DeepEqual(f.Get(), expected) {
		t.Fatal("cannot set a 2D array and get the same value back.")
	}
}

func Test_Any2DPrepareForUpdates(t *testing.T) {
	f := Any2D{}

	values_i, values_p := generateTestData2D()
	f.Set(values_i)

	f.PrepareForUpdates(10, key.NewPKey(50), getTestOnsetFunc())

	verifyStashUpdateInfo(t, &f.Reserved)

	for i, p1 := range values_p {
		for j, p2 := range p1 {
			if !p2.prepped {
				t.Errorf("array element (%d, %d) was not prepared correctly", i, j)
			}
		}
	}

	f.Set(values_i)

	if !testOnsetCalled {
		t.Error("onset was not called")
	}
}

func Test_Any2DEgestValue(t *testing.T) {
	f := Any2D{}
	f.Set([][]primitive.Interface{
		{&TestPrimitive{s: "abc0"}, &TestPrimitive{s: "xyz0"}},
		{&TestPrimitive{s: "abc1"}, &TestPrimitive{s: "xyz1"}},
	})

	v := f.EgestValue()
	a, ok := v.([][]any)
	if !ok {
		t.Fatal("cannot convert value to [][]any")
	}
	if len(a) != 2 {
		t.Fatal("wrong number of elements returned.  Expecting 2 elements")
	}

	for i, row := range a {
		if len(row) != 2 {
			t.Fatal("wrong number of elements in row.  Expecting 2 elements")
		}

		m1, ok := row[0].(map[any]any)
		if !ok {
			t.Fatal("cannot convert element to map[any]any")
		}
		m1v, ok := m1["s"].(string)
		if !ok {
			t.Fatal("cannot convert element map item to string")
		}
		if m1v != fmt.Sprintf("abc%d", i) {
			t.Fatal("wrong string value for element")
		}
		m2, ok := row[1].(map[any]any)
		if !ok {
			t.Fatal("cannot convert element to map[any]any")
		}
		m2v, ok := m2["s"].(string)
		if !ok {
			t.Fatal("cannot convert element map item to string")
		}
		if m2v != fmt.Sprintf("xyz%d", i) {
			t.Fatal("wrong string value for element")
		}
	}
}

func createAny2DForTest() (*Any2D, []*TestPrimitive) {
	f := &Any2D{}
	p11 := &TestPrimitive{}
	p12 := &TestPrimitive{}
	p21 := &TestPrimitive{}
	p22 := &TestPrimitive{}
	a := [][]primitive.Interface{{p11, p12}, {p21, p22}}
	f.Set(a)
	return f, []*TestPrimitive{p11, p12, p21, p22}
}

func Test_Any2DIngestUpdate(t *testing.T) {

	f, tps := createAny2DForTest()

	m11 := map[any]any{"s": "Good"}
	m12 := map[any]any{"s": "Morning"}
	m21 := map[any]any{"s": "Guten"}
	m22 := map[any]any{"s": "Tag"}
	ma := [][]any{{m11, m12}, {m21, m22}}

	err := f.IngestValue(ma)
	if err != nil {
		t.Fatalf("unexpected error returned:  %s", err.Error())
	}

	if tps[0].s != "Good" {
		t.Fatal("primitive #1,1 not updated correctly")
	}

	if tps[1].s != "Morning" {
		t.Fatal("primitive #1,2 not updated correctly")
	}

	if tps[2].s != "Guten" {
		t.Fatal("primitive #2,1 not updated correctly")
	}

	if tps[3].s != "Tag" {
		t.Fatal("primitive #2,2 not updated correctly")
	}
}

func Test_Any2DIngestUpdateInvalid1(t *testing.T) {

	f, _ := createAny2DForTest()

	err := f.IngestValue(3453)
	if err == nil {
		t.Fatal("no error returned for an invalid update")
	}
	if err.Error() != "invalid update" {
		t.Fatalf("wrong error was returned:  %s", err.Error())
	}
}

func Test_Any2DIngestUpdateInvalid2(t *testing.T) {

	f, _ := createAny2DForTest()

	err := f.IngestValue([][]any{{"Hello", "World"}, {"Hello", "World"}})

	if err == nil {
		t.Fatal("no error returned for an invalid update")
	}
	if err.Error() != "invalid update" {
		t.Fatalf("wrong error was returned:  %s", err.Error())
	}
}

func Test_Any2DIngestUpdateInvalidNumRows(t *testing.T) {

	f, _ := createAny2DForTest()

	m1 := map[any]any{"s": "Hello"}
	m2 := map[any]any{"s": "Hello"}

	err := f.IngestValue([][]any{{m1, m2}})

	if err == nil {
		t.Fatal("no error returned for an invalid update")
	}
	if err.Error() != "number of primitives in update does not equal existing primitives" {
		t.Fatalf("wrong error was returned:  %s", err.Error())
	}
}

func Test_Any2DIngestUpdateInvalidNumCols(t *testing.T) {

	f, _ := createAny2DForTest()

	m1 := map[any]any{"s": "Hello"}
	m2 := map[any]any{"s": "Hello"}

	err := f.IngestValue([][]any{{m1}, {m2}})

	if err == nil {
		t.Fatal("no error returned for an invalid update")
	}
	if err.Error() != "number of primitives in update does not equal existing primitives" {
		t.Fatalf("wrong error was returned:  %s", err.Error())
	}
}
