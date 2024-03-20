package field

import (
	"reflect"
	"testing"

	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

type TestPrimitive struct {
	b bool
	s string
	i int
}

// Implement primitive interface
func (*TestPrimitive) PrepareForUpdates(key.PKey, key.OnSetFunction) {
}

func Test_Any1DSetAndGet(t *testing.T) {
	f := Any1D{}

	item1 := &TestPrimitive{false, "a", 0}
	item2 := &TestPrimitive{true, "b", 1}

	aa := []primitive.Interface{item1, item2}

	f.Set(aa)

	if !reflect.DeepEqual(f.Get(), aa) {
		t.Fatal("cannot set any array and get the same value back.")
	}
}
