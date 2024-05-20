package field

import (
	"errors"
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

func (tp *TestPrimitive) LocateNextDescendant(locator *key.PKeyLocator) primitive.Interface {
	return nil
}

func (tp *TestPrimitive) EgestUpdate(fullupdate bool, fkeys []key.FKey) map[any]any {
	return map[any]any{"s": tp.s}
}

func (tp *TestPrimitive) IngestUpdate(update map[any]any) error {
	v, ok := update["s"]
	if !ok {
		return errors.New("field s not found in update")
	}
	tp.s, ok = v.(string)
	if !ok {
		return errors.New("field s in update cannot be converted to a string")
	}
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

	if f.fkey != 10 {
		t.Error("fkey was not stashed correctly")
	}
	if !f.pkey.EqualTo(key.NewPKey(50)) {
		t.Error("pkey was not stashed correctly")
	}
	if f.onset == nil {
		t.Error("onset was not stashed correctly")
	}
}

// Using a method to get a test function for onset in order to insure test state is reset in between
// tests.  Otherwise, the command-line tests will behave differently than those
// run in the IDE.
func getTestOnsetFunc() key.OnSetFunction {
	testOnsetCalled = false
	return _testOnset
}

var testOnsetCalled = false

func _testOnset(key.PKey, key.FKey, bool) {
	testOnsetCalled = true
}

func verifyIngestUpdateSuccessful(t *testing.T, err error, testfunc func() bool) {

	if err != nil {
		t.Fatalf("ingesting update returned error:  %s", err.Error())
	}

	if !testfunc() {
		t.Error("update not ingested correctly.  Expecting field value to be set correctly")
	}

	if testOnsetCalled {
		t.Error("onset was unexpectedly called while injesting update")
	}

}

func verifyIngestUpdateInvalid(t *testing.T, err error) {
	if err == nil {
		t.Fatal("no error returned after attemping to ingest invalid field value")
	}
	if err.Error() != "unable to convert value (any) to field value" {
		t.Fatal("wrong error was returned after attemping to ingest invalid field value")
	}
}
