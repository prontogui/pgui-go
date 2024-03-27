package golib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/prontogui/golib/primitive"
	// "github.com/prontogui/golib/field"
)

func trueOrFalseVariation(variation int) bool {
	if variation%1 == 0 {
		return true
	} else {
		return false
	}
}

func buildSimplePrimitiveArray(variation int) []primitive.Interface {

	p1 := &SimplePrimitive{}

	p1.Issued.Set(trueOrFalseVariation(variation))
	p1.Status.Set(100 + variation)
	p1.Embodiment.Set(fmt.Sprintf("apple%d", variation))

	p2 := &SimplePrimitive{}
	p2.Issued.Set(trueOrFalseVariation(variation))
	p2.Status.Set(150 + variation)
	p2.Embodiment.Set(fmt.Sprintf("orange%d", variation))

	return []primitive.Interface{p1, p2}
}

func buildSimplePrimitiveArray2D() [][]primitive.Interface {
	return [][]primitive.Interface{buildSimplePrimitiveArray(0), buildSimplePrimitiveArray(1)}
}

func verifyValueB(t *testing.T, value any, checkno int, expecting bool) {
	b, ok := value.(bool)
	if !ok {
		t.Errorf("for check #%d, unable to convert value to bool", checkno)
	}
	if b != expecting {
		t.Errorf("for check #%d, value is %v.  Expecting %v", checkno, b, expecting)
	}
}

func verifyValueI(t *testing.T, value any, checkno int, expecting int) {
	i, ok := value.(int)
	if !ok {
		t.Errorf("for check #%d, unable to convert value to int", checkno)
	}
	if i != expecting {
		t.Errorf("for check #%d, value is %v.  Expecting %v", checkno, i, expecting)
	}
}

func verifyValueS(t *testing.T, value any, checkno int, expecting string) {
	s, ok := value.(string)
	if !ok {
		t.Errorf("for check #%d, unable to convert value to string", checkno)
	}
	if s != expecting {
		t.Errorf("for check #%d, value is %v.  Expecting %v", checkno, s, expecting)
	}
}

func verifyValueSA(t *testing.T, value any, checkno int, expecting []string) {
	sa, ok := value.([]string)
	if !ok {
		t.Errorf("for check #%d, unable to convert value to []string", checkno)
	}
	if !reflect.DeepEqual(sa, expecting) {
		t.Errorf("for check #%d, value is %v.  Expecting %v", checkno, sa, expecting)
	}
}

func verifyValueBL(t *testing.T, value any, checkno int, expecting []byte) {
	ba, ok := value.([]byte)
	if !ok {
		t.Errorf("for check #%d, unable to convert value to []byte", checkno)
	}
	if !reflect.DeepEqual(ba, expecting) {
		t.Errorf("for check #%d, value is %v.  Expecting %v", checkno, ba, expecting)
	}
}

func verifySimpleElement(t *testing.T, v any, checkno int, b bool, i int, s string) {
	m, ok := v.(map[string]any)
	if !ok {
		t.Errorf("for check #%d, unable to convert item %d to map[string]any", checkno, i)
	}
	verifyValueB(t, m["Issued"], checkno, b)
	verifyValueI(t, m["Status"], checkno, i)
	verifyValueS(t, m["Embodiment"], checkno, s)
}

func verifyValueA1D(t *testing.T, value any, checkno int, variation int) {
	a1d, ok := value.([]any)
	if !ok {
		t.Errorf("for check #%d, unable to convert value to []any", checkno)
	}

	verifySimpleElement(t, a1d[0], checkno, trueOrFalseVariation(variation), 100+variation, fmt.Sprintf("apple%d", variation))
	verifySimpleElement(t, a1d[1], checkno, trueOrFalseVariation(variation), 150+variation, fmt.Sprintf("orange%d", variation))
}

func verifyValueA2D(t *testing.T, value any, checkno int) {
	a2d, ok := value.([][]any)
	if !ok {
		t.Errorf("for check #%d, unable to convert value to [][]any", checkno)
	}

	verifyValueA1D(t, a2d[0], checkno, 0)
	verifyValueA1D(t, a2d[1], checkno, 1)
}

func Test_EgestUpdate(t *testing.T) {

	tp := ComplexPrimitive{}
	tp.Issued.Set(true)
	tp.Status.Set(100)
	tp.Embodiment.Set("industry")
	tp.Choices.Set([]string{"abc", "def", "xyz"})
	tp.Data.Set([]byte{100, 150, 200})
	tp.ListItems.Set(buildSimplePrimitiveArray(0))
	tp.Rows.Set(buildSimplePrimitiveArray2D())

	tp.PrepareForUpdates(0, nil)

	update := tp.EgestUpdate(true, nil)

	if update == nil {
		t.Fatal("returned nil.  Expecting a valid map")
	}

	updatelen := len(update)
	if updatelen != 7 {
		t.Fatalf("returned map has %d items.  Expecting 7 items", updatelen)
	}

	verifyValueB(t, update["Issued"], 1, true)
	verifyValueI(t, update["Status"], 2, 100)
	verifyValueS(t, update["Embodiment"], 3, "industry")
	verifyValueSA(t, update["Choices"], 4, []string{"abc", "def", "xyz"})
	verifyValueA1D(t, update["ListItems"], 5, 0)
	verifyValueA2D(t, update["Rows"], 6)
	verifyValueBL(t, update["Data"], 7, []byte{100, 150, 200})
}

func Test_IngestUpdate(t *testing.T) {

	choices := []string{"A", "B", "C"}
	a1d0 := map[string]any{"Issued": false, "Embodiment": "fabricated"}
	a1d1 := map[string]any{"Issued": true, "Embodiment": "made up"}

	update := map[string]any{
		"Issued":     true,
		"Status":     99,
		"Embodiment": "apple",
		"Choices":    choices,
		"ListItems":  []any{a1d0, a1d1},
	}

	tp := ComplexPrimitive{}
	tp.PrepareForUpdates(0, nil)

	err := tp.IngestUpdate(update)
	if err != nil {
		t.Fatalf("unexpected error returned:  %s", err.Error())
	}

	if tp.Issued.Get() != true {
		t.Error("field Issued was not updated correctly")
	}
	if tp.Status.Get() != 99 {
		t.Error("field Status was not updated correctly")
	}
	if tp.Embodiment.Get() != "apple" {
		t.Error("field Embodiment was not updated correctly")
	}
	if !reflect.DeepEqual(tp.Choices.Get(), choices) {
		t.Error("field Choices was not updated correctly")
	}
	if len(tp.ListItems.Get()) != 2 {
		t.Fatal("field ListItems was not updated correctly")
	}
	p1 := tp.ListItems.Get()[0].(*ComplexPrimitive)
	p2 := tp.ListItems.Get()[1].(*ComplexPrimitive)
	if p1.Issued.Get() != false || p1.Embodiment.Get() != "fabricated" {
		t.Error("element 0 of field ListItems was not updated correctly")
	}
	if p2.Issued.Get() != true || p2.Embodiment.Get() != "made up" {
		t.Error("element 1 of field ListItems was not updated correctly")
	}
}

func Test_IngestUpdateInvalidFieldName(t *testing.T) {

	update := map[string]any{
		"ASDFLKHMN2KJESRHFNASDFASDFGCVC": true,
	}

	tp := ComplexPrimitive{}
	tp.PrepareForUpdates(0, nil)

	err := tp.IngestUpdate(update)
	if err == nil {
		t.Fatal("no error returned.  Expected an error since update specifies a field that doesn't exist")
	}
	if err.Error() != "invalid field name" {
		t.Fatal("wrong error was returned")
	}
}

func Test_IngestUpdateNoMatchingFieldInPrimitive(t *testing.T) {

	update := map[string]any{
		"Choices": []string{},
	}

	tp := SimplePrimitive{}
	tp.PrepareForUpdates(0, nil)

	err := tp.IngestUpdate(update)
	if err == nil {
		t.Fatal("no error returned.  Expected an error since update specifies a field that doesn't exist in primitive")
	}
	if err.Error() != "no matching field name in primitive" {
		t.Fatal("wrong error was returned")
	}
}
