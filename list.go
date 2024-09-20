// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A list is a collection of primitives that have a sequential-like relationship
// and might be dynamic in quantity or kind.
type ListWith struct {
	Embodiment   string
	ListItems    []Primitive
	Selected     int
	TemplateItem Primitive
}

// Creates a new List using the supplied field assignments.
func (w ListWith) Make() *List {
	list := &List{}
	list.SetEmbodiment(w.Embodiment)
	list.listItems.Set(w.ListItems)
	list.SetSelected(w.Selected)
	list.SetTemplateItem(w.TemplateItem)
	return list
}

// A list is a collection of primitives that have a sequential-like relationship
// and might be dynamic in quantity or kind.
type List struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment   StringField
	listItems    Any1DField
	selected     IntegerField
	templateItem AnyField
}

// Creates a new List and assigns items.
func NewList(items ...Primitive) *List {
	return ListWith{ListItems: items}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (list *List) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	list.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &list.embodiment},
			{key.FKey_ListItems, &list.listItems},
			{key.FKey_Selected, &list.selected},
			{key.FKey_TemplateItem, &list.templateItem},
		}
	})
}

// A non-recursive method to locate descendants by PKey.  This is used internally by this library
// and normally should not be called by users of the library.
func (list *List) LocateNextDescendant(locator *key.PKeyLocator) Primitive {
	// TODO:  generalize this code by handling inside primitive Reserved area.

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

// Returns a JSON string specifying the embodiment to use for this primitive.
func (list *List) Embodiment() string {
	return list.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (list *List) SetEmbodiment(s string) *List {
	list.embodiment.Set(s)
	return list
}

// Returns the ttems to show in the list.
func (list *List) ListItems() []Primitive {
	return list.listItems.Get()
}

// Sets the ttems to show in the list.
func (list *List) SetListItems(items []Primitive) *List {
	list.listItems.Set(items)
	return list
}

// Returns the ttems to show in the list (as a variadic argument list).
func (list *List) SetListItemsVA(items ...Primitive) *List {
	list.listItems.Set(items)
	return list
}

// Returns the currently selected item or -1 for none selected.
func (list *List) Selected() int {
	return list.selected.Get()
}

// Sets the currently selected item or -1 for none selected.
func (list *List) SetSelected(selected int) *List {
	list.selected.Set(selected)
	return list
}

// Returns the template for how each item in the list should look, feel, and behave.
func (list *List) TemplateItem() Primitive {
	return list.templateItem.Get()
}

// Sets the template for how each item in the list should look, feel, and behave.
func (list *List) SetTemplateItem(item Primitive) *List {
	list.templateItem.Set(item)
	return list
}
