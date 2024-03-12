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

	updateIn, err := ExchangeUpdates("12")

	if err != nil {
		t.Fatal("error was returned.  Expected no error")
	}
	if updateIn.(string) != "12" {
		t.Fatal("wrong update was returned")
	}
}

// Test proper handling of the inboundUpdates channel being closed during an exchange.
func Test_ExchangeUpdates2(t *testing.T) {

	go func() {
		<-outboundUpdates
		close(inboundUpdates)
	}()

	_, err := ExchangeUpdates("12")

	if err == nil {
		t.Fatal("no error was returned.  Expected an error")
	}

	if err.Error() != "inboundUpdates channel is invalid" {
		t.Fatal("wrong error was returned")
	}
}
