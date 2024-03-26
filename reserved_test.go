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
	if variation%1 == 0 {
		return true
	} else {
		return false
	}
}
func buildSimplePrimitiveArray(variation int) []primitive.Interface {

	p1 := &SimplePrimitive{}

	p1.B.Set(trueOrFalseVariation(variation))
	p1.I.Set(100 + variation)
	p1.S.Set(fmt.Sprintf("apple%d", variation))

	p2 := &SimplePrimitive{}
	p2.B.Set(trueOrFalseVariation(variation))
	p2.I.Set(150 + variation)
	p2.S.Set(fmt.Sprintf("orange%d", variation))

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

func verifySimpleElement(t *testing.T, v any, checkno int, b bool, i int, s string) {
	m, ok := v.(map[string]any)
	if !ok {
		t.Errorf("for check #%d, unable to convert item %d to map[string]any", checkno, i)
	}
	verifyValueB(t, m["B"], checkno, b)
	verifyValueI(t, m["I"], checkno, i)
	verifyValueS(t, m["S"], checkno, s)
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
	tp.B.Set(true)
	tp.I.Set(100)
	tp.S.Set("industry")
	tp.SA.Set([]string{"abc", "def", "xyz"})
	tp.BL.Set([]byte{100, 150, 200})
	tp.A1D.Set(buildSimplePrimitiveArray(0))
	tp.A2D.Set(buildSimplePrimitiveArray2D())

	testOnset := func(key.PKey, key.FKey, bool) {

	}

	tp.PrepareForUpdates(0, testOnset)

	update := tp.EgestUpdate(true, nil)

	if update == nil {
		t.Fatal("returned nil.  Expecting a valid map")
	}

	updatelen := len(update)
	if updatelen != 7 {
		t.Fatalf("returned map has %d items.  Expecting 7 items", updatelen)
	}

	verifyValueB(t, update["B"], 1, false)
	verifyValueI(t, update["I"], 2, 101)
	verifyValueS(t, update["S"], 3, "world")
	verifyValueSA(t, update["SA"], 4, []string{"ab", "de", "xy"})
	verifyValueA1D(t, update["A1D"], 5, 0)
	verifyValueA2D(t, update["A2D"], 6)
}
