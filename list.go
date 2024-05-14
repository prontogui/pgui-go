package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

type ListWith struct {
	ListItems []primitive.Interface
	Selected  int
}

func (w ListWith) Make() *List {
	list := &List{}
	list.listItems.Set(w.ListItems)
	list.SetSelected(w.Selected)
	return list
}

type List struct {
	// Mix-in the common guts for primitives (B-side fields, implements primitive interface, etc.)
	Reserved

	listItems field.Any1D
	selected  field.Integer
}

func (list *List) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	list.AttachField("ListItems", &list.listItems)
	list.AttachField("Selected", &list.selected)

	// Prepare all fields for updates
	list.Reserved.PrepareForUpdates(pkey, onset)
}

func (list *List) GetChildPrimitive(index int) primitive.Interface {
	listItems := list.listItems.Get()

	if index < len(listItems) {
		return listItems[index]
	}

	return nil
}

func (list *List) ListItems() []primitive.Interface {
	return list.listItems.Get()
}

func (list *List) SetListItems(items []primitive.Interface) {
	list.listItems.Set(items)
}

func (list *List) SetListItemsVA(items ...primitive.Interface) {
	list.listItems.Set(items)
}

func (list *List) Selected() int {
	return -1
}

func (list *List) SetSelected(selected int) {

}
