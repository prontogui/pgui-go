package field

import (
	"testing"
)

func Test_StringSetAndGet(t *testing.T) {
	f := String{}

	f.Set("abc")

	if f.Get() != "abc" {
		t.Fatal("cannot set string and get the same value back.")
	}
}
