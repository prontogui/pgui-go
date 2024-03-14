package field

import (
	"testing"
)

func Test_IntegerSetAndGet(t *testing.T) {
	f := Integer{}

	f.Set(92342)

	if f.Get() != 92342 {
		t.Fatal("cannot set integer and get the same value back.")
	}
}
