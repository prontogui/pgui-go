package field

import (
	"reflect"
	"testing"
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

func Test_Any1DIngestUpdate(t *testing.T) {

	f := Any1D{}
	err := f.IngestUpdate([]any{})
	if err == nil || err.Error() != "ingesting field update for Any1D is not supported" {
		t.Fatal("ingesting update for Any1D should not be supported yet")
	}
}
