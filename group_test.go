package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

func Test_GroupAttachedFields(t *testing.T) {
	grp := &Group{}
	grp.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, grp.Reserved, "GroupItems")
}

func Test_GroupMake(t *testing.T) {
	grp := GroupWith{GroupItems: []primitive.Interface{&Command{}, &Command{}}}.Make()

	if len(grp.GroupItems()) != 2 {
		t.Error("'GroupItems' field was not initialized correctly")
	}
}

func Test_GruopFieldSettings(t *testing.T) {

	grp := &Group{}

	grp.SetGroupItems([]primitive.Interface{&Command{}, &Command{}})

	grpGet := grp.GroupItems()

	if len(grpGet) != 2 {
		t.Errorf("GroupItems() returned %d items.  Expecting 2 items.", len(grpGet))
	}

	_, ok1 := grpGet[0].(*Command)
	if !ok1 {
		t.Error("First group is not a Command primitive.")
	}
	_, ok2 := grpGet[1].(*Command)
	if !ok2 {
		t.Error("Second group is not a Command primitive.")
	}

	grp.SetGroupItemsVA(&Text{}, &Text{})

	grpGet = grp.GroupItems()

	if len(grpGet) != 2 {
		t.Errorf("GroupItems() returned %d items after calling variadic setter.  Expecting 2 items.", len(grpGet))
	}

	_, ok1 = grpGet[0].(*Text)
	if !ok1 {
		t.Error("First group is not a Text primitive.")
	}
	_, ok2 = grpGet[1].(*Text)
	if !ok2 {
		t.Error("Second group is not a Text primitive.")
	}

}
