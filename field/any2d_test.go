package field

import (
	"fmt"
	"reflect"
	"testing"

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

	f.PrepareForUpdates(10, 50, getTestOnsetFunc())

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

		m1, ok := row[0].(map[string]any)
		if !ok {
			t.Fatal("cannot convert element to map[string]any")
		}
		m1v, ok := m1["s"].(string)
		if !ok {
			t.Fatal("cannot convert element map item to string")
		}
		if m1v != fmt.Sprintf("abc%d", i) {
			t.Fatal("wrong string value for element")
		}
		m2, ok := row[1].(map[string]any)
		if !ok {
			t.Fatal("cannot convert element to map[string]any")
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

func Test_Any2DIngestUpdate(t *testing.T) {

	f := Any2D{}
	err := f.IngestValue([][]any{})
	if err == nil || err.Error() != "ingesting value for Any2D is not supported" {
		t.Fatal("ingesting value for Any2D should not be supported yet")
	}
}
