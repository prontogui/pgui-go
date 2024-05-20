package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

type ListWith struct {
	ListItems      []primitive.Interface
	Selected       int
	ItemEmbodiment string
}

func (w ListWith) Make() *List {
	list := &List{}
	list.listItems.Set(w.ListItems)
	list.SetSelected(w.Selected)
	list.SetItemEmbodiment(w.ItemEmbodiment)
	return list
}

type List struct {
	// Mix-in the common guts for primitives (B-side fields, implements primitive interface, etc.)
	Reserved

	listItems      field.Any1D
	selected       field.Integer
	itemEmbodiment field.String
}

func (list *List) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	list.AttachField("ListItems", &list.listItems)
	list.AttachField("Selected", &list.selected)
	list.AttachField("ItemEmbodiment", &list.itemEmbodiment)

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
	return list.selected.Get()
}

func (list *List) SetSelected(selected int) {
	list.selected.Set(selected)
}

func (list *List) ItemEmbodiment() string {
	return list.itemEmbodiment.Get()
}

func (list *List) SetItemEmbodiment(embodiment string) {
	list.itemEmbodiment.Set(embodiment)
}
