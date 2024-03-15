package field

import (
	"reflect"
	"testing"
)

func Test_Any1DSetAndGet(t *testing.T) {
	f := Any1D{}

	aa := []any{true, "abc", 10}

	f.Set(aa)

	if !reflect.DeepEqual(f.Get(), aa) {
		t.Fatal("cannot set any array and get the same value back.")
	}
}
