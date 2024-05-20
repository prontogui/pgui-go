package golib

import (
	"reflect"
	"testing"

	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

func Test_ListAttachedFields(t *testing.T) {
	list := &List{}
	list.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, list.Reserved, "ListItems", "Selected", "TemplateItem")
}

func Test_ListMake(t *testing.T) {
	list := ListWith{
		ListItems:    []primitive.Interface{&Command{}, &Command{}},
		Selected:     1,
		TemplateItem: &Command{},
	}.Make()

	if len(list.ListItems()) != 2 {
		t.Error("'ListItems' field was not initialized correctly")
	}

	if list.Selected() != 1 {
		t.Error("List selection not initialized properly")
	}

	if !reflect.DeepEqual(list.TemplateItem(), &Command{}) {
		t.Error("TemplateItem is not initialized properly")
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

	// TemplateItem field tests
	list.SetTemplateItem(&Text{})
	if !reflect.DeepEqual(list.TemplateItem(), &Text{}) {
		t.Error("Unable to set template item to a Text primitive")
	}
}

func Test_ListGetChildPrimitive(t *testing.T) {

	list := &List{}

	cmd1 := CommandWith{Label: "a"}.Make()
	cmd2 := CommandWith{Label: "b"}.Make()
	cmd3 := CommandWith{Label: "c"}.Make()

	list.SetListItemsVA(cmd1, cmd2)
	list.SetTemplateItem(cmd3)

	locate := func(pkey key.PKey) *Command {
		locator := key.NewPKeyLocator(pkey)
		return list.LocateNextDescendant(locator).(*Command)
	}

	if locate(key.NewPKey(0, 0)).Label() != "a" {
		t.Fatal("LocateNextDescendant doesn't return a child for pkey 0, 0.")
	}

	if locate(key.NewPKey(0, 1)).Label() != "b" {
		t.Fatal("LocateNextDescendant doesn't return a child for pkey 0, 1.")
	}

	if locate(key.NewPKey(1)).Label() != "c" {
		t.Fatal("LocateNextDescendant doesn't return a child for pkey 1. ")
	}
}
