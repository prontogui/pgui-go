package field

import (
	"reflect"
	"testing"
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

func Test_Any2DIngestUpdate(t *testing.T) {

	f := Any2D{}
	err := f.IngestValue([][]any{})
	if err == nil || err.Error() != "ingesting value for Any2D is not supported" {
		t.Fatal("ingesting value for Any2D should not be supported yet")
	}
}
