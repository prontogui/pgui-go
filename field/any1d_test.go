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

func Test_Any1DIngestUpdate(t *testing.T) {

	f := Any1D{}
	err := f.IngestValue([]any{})
	if err == nil || err.Error() != "ingesting value for Any1D is not supported" {
		t.Fatal("ingesting value for Any1D should not be supported yet")
	}
}
