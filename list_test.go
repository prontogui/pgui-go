package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

func Test_ListAttachedFields(t *testing.T) {
	list := &List{}
	list.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, list.Reserved, "ListItems", "Selected", "ItemEmbodiment")
}

func Test_ListMake(t *testing.T) {
	list := ListWith{
		ListItems:      []primitive.Interface{&Command{}, &Command{}},
		Selected:       1,
		ItemEmbodiment: "some-embodiment",
	}.Make()

	if len(list.ListItems()) != 2 {
		t.Error("'ListItems' field was not initialized correctly")
	}

	if list.Selected() != 1 {
		t.Error("List selection not initialized properly")
	}

	if list.ItemEmbodiment() != "some-embodiment" {
		t.Error("ItemEmbodiment is not initialized properly")
	}
}

func Test_ListFieldSettings(t *testing.T) {

	list := &List{}

	// ListItems field (as array)

	list.SetListItems([]primitive.Interface{&Command{}, &Command{}})

	listGet := list.ListItems()

	if len(listGet) != 2 {
		t.Errorf("ListItems() returned %d items.  Expecting 2 items.", len(listGet))
	}

	_, ok1 := listGet[0].(*Command)
	if !ok1 {
		t.Error("First group is not a Command primitive.")
	}
	_, ok2 := listGet[1].(*Command)
	if !ok2 {
		t.Error("Second group is not a Command primitive.")
	}

	// ListItems field (as variadic items)

	list.SetListItemsVA(&Text{}, &Text{})

	listGet = list.ListItems()

	if len(listGet) != 2 {
		t.Errorf("ListItems() returned %d items after calling variadic setter.  Expecting 2 items.", len(listGet))
	}

	_, ok1 = listGet[0].(*Text)
	if !ok1 {
		t.Error("First item is not a Text primitive.")
	}
	_, ok2 = listGet[1].(*Text)
	if !ok2 {
		t.Error("Second item is not a Text primitive.")
	}

	// Selected field tests

	list.SetSelected(-1)
	if list.Selected() != -1 {
		t.Error("Unable to set seletion to -1")
	}

	list.SetSelected(0)
	if list.Selected() != 0 {
		t.Error("Unable to set seletion to 0")
	}

	list.SetSelected(1)
	if list.Selected() != 1 {
		t.Error("Unable to set seletion to 1")
	}

	// ItemEmbodiment field tests
	list.SetItemEmbodiment("some-thing")
	if list.ItemEmbodiment() != "some-thing" {
		t.Error("Unable to set item embodiment to 'some-thing'")
	}
}

func Test_ListGetChildPrimitive(t *testing.T) {

	list := &List{}

	list.SetListItems([]primitive.Interface{&Command{}, &Command{}})

	if list.GetChildPrimitive(0) == nil {
		t.Fatal("GetChildPrimitve doesn't return a child for index = 0.")
	}

	if list.GetChildPrimitive(1) == nil {
		t.Fatal("GetChildPrimitve doesn't return a child for index = 1.")
	}

	if list.GetChildPrimitive(2) != nil {
		t.Fatal("GetChildPrimitve shouldn't return a child for index = 2.")
	}
}
