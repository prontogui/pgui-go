package field

import (
	"testing"

	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

type TestPrimitive struct {
	s       string
	prepped bool
}

func (tp *TestPrimitive) PrepareForUpdates(key.PKey, key.OnSetFunction) {
	tp.prepped = true
}

func (tp *TestPrimitive) GetChildPrimitive(index int) primitive.Interface {
	return nil
}

func (tp *TestPrimitive) GetFieldValue(fieldname string) any {
	return nil
}

func generateTestData1D() ([]primitive.Interface, []*TestPrimitive) {

	act1 := &TestPrimitive{s: "abc"}
	act2 := &TestPrimitive{s: "def"}
	act3 := &TestPrimitive{s: "uvw"}

	return []primitive.Interface{act1, act2, act3}, []*TestPrimitive{act1, act2, act3}
}

func generateTestData2D() ([][]primitive.Interface, [][]*TestPrimitive) {

	act1a := &TestPrimitive{s: "abc"}
	act1b := &TestPrimitive{s: "def"}
	act2a := &TestPrimitive{s: "uvw"}
	act2b := &TestPrimitive{s: "xyz"}

	return [][]primitive.Interface{{act1a, act1b}, {act2a, act2b}}, [][]*TestPrimitive{{act1a, act1b}, {act2a, act2b}}
}

func verifyStashUpdateInfo(t *testing.T, f *Reserved) {

	if f.fkey != key.FKeyFor("Abc") {
		t.Error("fkey was not stashed correctly")
	}
	if f.pkey != 50 {
		t.Error("pkey was not stashed correctly")
	}
	if f.onset == nil {
		t.Error("onset was not stashed correctly")
	}
}

var testOnsetCalled = false

func testOnset(key.PKey, key.FKey, bool) {
	testOnsetCalled = true
}
