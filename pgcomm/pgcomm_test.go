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

// Test the normal exchange of updates between server and the app.
func Test_ExchangeUpdates1(t *testing.T) {

	go func() {
		update, _ := <-outboundUpdates
		inboundUpdates <- update
	}()

	ok, updateIn := ExchangeUpdates([]byte{1, 2})

	if !ok {
		t.Fatal("ok = false was returned.  Expected ok = true")
	}
	if updateIn == nil || len(updateIn) != 2 || updateIn[0] != 1 || updateIn[1] != 2 {
		t.Fatal("wrong update was returned")
	}
}

// Test proper handling of the inboundUpdates channel being closed during an exchange.
func Test_ExchangeUpdates2(t *testing.T) {

	go func() {
		<-outboundUpdates
		close(inboundUpdates)
	}()

	ok, _ := ExchangeUpdates([]byte{1, 2})

	if ok {
		t.Fatal("ok = true was returned.  Expected ok = false")
	}
}
