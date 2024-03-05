package pgcomm

import (
	"testing"

	"github.com/prontogui/golib/testhelp"
)

func Test_serve_badport(t *testing.T) {
	err := StartServing("", -1)
	testhelp.TestErrorMessage(t, err, "listen tcp: address -1: invalid port")
}

func Test_serve_good(t *testing.T) {
	err := StartServing("", 0)
	testhelp.TestNilError(t, err)
	StopServing()
}
