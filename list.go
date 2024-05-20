package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

type ListWith struct {
	ListItems    []primitive.Interface
	Selected     int
	TemplateItem primitive.Interface
}

func (w ListWith) Make() *List {
	list := &List{}
	list.listItems.Set(w.ListItems)
	list.SetSelected(w.Selected)
	list.SetTemplateItem(w.TemplateItem)
	return list
}

type List struct {
	// Mix-in the common guts for primitives (B-side fields, implements primitive interface, etc.)
	Reserved

	listItems    field.Any1D
	selected     field.Integer
	templateItem field.Any
}

func (list *List) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	list.AttachField("ListItems", &list.listItems)
	list.AttachField("Selected", &list.selected)
	list.AttachField("TemplateItem", &list.templateItem)

	// Prepare all fields for updates
	list.Reserved.PrepareForUpdates(pkey, onset)
}

func (list *List) LocateNextDescendant(locator *key.PKeyLocator) primitive.Interface {

	nextIndex := locator.NextIndex()

	// Fields are handled in alphabetical order
	switch nextIndex {
	case 0:
		return list.ListItems()[locator.NextIndex()]
	case 1:
		return list.TemplateItem()
	default:
		panic("cannot locate descendent using a pkey that we assumed was valid")
	}
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

func (list *List) TemplateItem() primitive.Interface {
	return list.templateItem.Get()
}

func (list *List) SetTemplateItem(item primitive.Interface) {
	list.templateItem.Set(item)
}
