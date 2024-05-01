package golib

import (
	"reflect"
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_ChoiceAttachedFields(t *testing.T) {
	cmd := &Choice{}
	cmd.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, cmd.Reserved, "Choice", "Issued", "Choices")
}

func Test_ChoiceFieldSettings(t *testing.T) {
	choice := &Choice{}

	choice.SetChoice("Big Fish")
	if choice.Choice() != "Big Fish" {
		t.Error("Could not set Choice field.")
	}

	choice.changed.Set(true)
	if !choice.Changed() {
		t.Error("Could not get Changed field correctly.")
	}

	choice.SetChoices([]string{"mary", "john", "paul"})
	if !reflect.DeepEqual(choice.Choices(), []string{"mary", "john", "paul"}) {
		t.Error("Could not set Choices field.")
	}

	choice.SetChoicesVA("nancy", "tom", "bob")
	if !reflect.DeepEqual(choice.Choices(), []string{"nancy", "tom", "bob"}) {
		t.Error("Could not set Choices field using variadic arguments.")
	}

}
