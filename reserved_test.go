package golib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
	// "github.com/prontogui/golib/field"
)

func trueOrFalseVariation(variation int) bool {
	if (variation % 2) == 0 {
		return true
	} else {
		return false
	}
}

func buildSimplePrimitiveArray(variation int) []primitive.Interface {

	p1 := &SimplePrimitive{}

	p1.Issued.Set(trueOrFalseVariation(variation))
	p1.Status.Set(100 + variation)
	p1.Label.Set(fmt.Sprintf("apple%d", variation))

	p2 := &SimplePrimitive{}
	p2.Issued.Set(trueOrFalseVariation(variation))
	p2.Status.Set(150 + variation)
	p2.Label.Set(fmt.Sprintf("orange%d", variation))

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
	m, ok := v.(map[any]any)
	if !ok {
		t.Errorf("for check #%d, unable to convert item %d to map[any]any", checkno, i)
	}
	verifyValueB(t, m["Issued"], checkno, b)
	verifyValueI(t, m["Status"], checkno, i)
	verifyValueS(t, m["Label"], checkno, s)
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

func createEgestPrimitiveForTest() *ComplexPrimitive {
	tp := &ComplexPrimitive{}
	tp.Issued.Set(true)
	tp.Status.Set(100)
	tp.Choices.Set([]string{"abc", "def", "xyz"})
	tp.Data.Set([]byte{100, 150, 200})
	tp.ListItems.Set(buildSimplePrimitiveArray(0))
	tp.Rows.Set(buildSimplePrimitiveArray2D())
	return tp
}

func Test_EgestFullUpdate(t *testing.T) {

	tp := createEgestPrimitiveForTest()

	tp.PrepareForUpdates(key.NewPKey(0), nil)

	update := tp.EgestUpdate(true, nil)

	if update == nil {
		t.Fatal("returned nil.  Expecting a valid map")
	}

	updatelen := len(update)
	if updatelen != 6 {
		t.Fatalf("returned map has %d items.  Expecting 6 items", updatelen)
	}

	verifyValueB(t, update["Issued"], 1, true)
	verifyValueI(t, update["Status"], 2, 100)
	verifyValueSA(t, update["Choices"], 6, []string{"abc", "def", "xyz"})
	verifyValueA1D(t, update["ListItems"], 7, 0)
	verifyValueA2D(t, update["Rows"], 8)
	verifyValueBL(t, update["Data"], 9, []byte{100, 150, 200})
}

func Test_EgestPartialUpdate(t *testing.T) {

	tp := createEgestPrimitiveForTest()

	tp.PrepareForUpdates(key.NewPKey(0), nil)

	tp.Issued.Set(true)
	tp.Status.Set(2)

	update := tp.EgestUpdate(false, []key.FKey{key.FKeyFor("Issued"), key.FKeyFor("Status")})

	if update == nil {
		t.Fatal("returned nil.  Expecting a valid map")
	}

	updatelen := len(update)
	if updatelen != 2 {
		t.Fatalf("returned map has %d items.  Expecting 2 items", updatelen)
	}

	verifyValueB(t, update["Issued"], 1, true)
	verifyValueI(t, update["Status"], 2, 2)
}

func Test_IngestUpdate(t *testing.T) {

	choices := []string{"A", "B", "C"}
	m1 := map[any]any{"Issued": false, "Label": "fabricated"}
	m2 := map[any]any{"Issued": true, "Label": "made up"}
	m11 := map[any]any{"Issued": true, "Label": "contrived"}
	m12 := map[any]any{"Issued": false, "Label": "imagined"}
	m21 := map[any]any{"Issued": false, "Label": "brainstormed"}
	m22 := map[any]any{"Issued": true, "Label": "revealed"}

	update := map[any]any{
		"Issued":    true,
		"Status":    99,
		"Choices":   choices,
		"ListItems": []any{m1, m2},
		"Rows":      [][]any{{m11, m12}, {m21, m22}},
	}

	tp := ComplexPrimitive{}
	p1 := &SimplePrimitive{}
	p2 := &SimplePrimitive{}
	p11 := &SimplePrimitive{}
	p12 := &SimplePrimitive{}
	p21 := &SimplePrimitive{}
	p22 := &SimplePrimitive{}

	tp.ListItems.Set([]primitive.Interface{p1, p2})
	tp.Rows.Set([][]primitive.Interface{{p11, p12}, {p21, p22}})
	tp.PrepareForUpdates(key.NewPKey(0), nil)

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

	if !reflect.DeepEqual(tp.Choices.Get(), choices) {
		t.Error("field Choices was not updated correctly")
	}
	if len(tp.ListItems.Get()) != 2 {
		t.Fatal("field ListItems was not updated correctly")
	}

	verifyPrimitiveElement := func(p *SimplePrimitive, desc string, issued bool, label string) {

		if p.Issued.Get() != issued || p.Label.Get() != label {
			t.Errorf("%s was not updated correctly", desc)
		}
	}

	verifyPrimitiveElement(p1, "element #1 of ListItems", false, "fabricated")
	verifyPrimitiveElement(p2, "element #2 of ListItems", true, "made up")
	verifyPrimitiveElement(p11, "element #1,1 of Rows", true, "contrived")
	verifyPrimitiveElement(p12, "element #1,2 of Rows", false, "imagined")
	verifyPrimitiveElement(p21, "element #2,1 of Rows", false, "brainstormed")
	verifyPrimitiveElement(p22, "element #2,2 of Rows", true, "revealed")

}

func Test_IngestUpdateInvalidFieldName(t *testing.T) {

	update := map[any]any{
		"ASDFLKHMN2KJESRHFNASDFASDFGCVC": true,
	}

	tp := ComplexPrimitive{}
	tp.PrepareForUpdates(key.NewPKey(0), nil)

	err := tp.IngestUpdate(update)
	if err == nil {
		t.Fatal("no error returned.  Expected an error since update specifies a field that doesn't exist")
	}
	if err.Error() != "invalid field name" {
		t.Fatal("wrong error was returned")
	}
}

func Test_IngestUpdateNoMatchingFieldInPrimitive(t *testing.T) {

	update := map[any]any{
		"Choices": []string{},
	}

	tp := SimplePrimitive{}
	tp.PrepareForUpdates(key.NewPKey(0), nil)

	err := tp.IngestUpdate(update)
	if err == nil {
		t.Fatal("no error returned.  Expected an error since update specifies a field that doesn't exist in primitive")
	}
	if err.Error() != "no matching field name in primitive" {
		t.Fatal("wrong error was returned")
	}
}
