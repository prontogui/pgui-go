package field

import (
	"reflect"
	"testing"
)

func Test_Any2DSetAndGet(t *testing.T) {
	f := Any2D{}

	aaa := [][]any{{true, "abc", 10}, {false, "def", 20}, {true, "xyz", 30}}

	f.Set(aaa)

	if !reflect.DeepEqual(f.Get(), aaa) {
		t.Fatal("cannot set any 2D array and get the same value back.")
	}
}
