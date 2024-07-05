// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

type ListWith struct {
	Embodiment   string
	ListItems    []Primitive
	Selected     int
	TemplateItem Primitive
}

func (w ListWith) Make() *List {
	list := &List{}
	list.SetEmbodiment(w.Embodiment)
	list.listItems.Set(w.ListItems)
	list.SetSelected(w.Selected)
	list.SetTemplateItem(w.TemplateItem)
	return list
}

type List struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment   StringField
	listItems    Any1DField
	selected     IntegerField
	templateItem AnyField
}

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

// TODO:  generalize this code by handling inside primitive Reserved area.
func (list *List) LocateNextDescendant(locator *key.PKeyLocator) Primitive {

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

func (list *List) Embodiment() string {
	return list.embodiment.Get()
}

func (list *List) SetEmbodiment(s string) {
	list.embodiment.Set(s)
}

func (list *List) ListItems() []Primitive {
	return list.listItems.Get()
}

func (list *List) SetListItems(items []Primitive) {
	list.listItems.Set(items)
}

func (list *List) SetListItemsVA(items ...Primitive) {
	list.listItems.Set(items)
}

func (list *List) Selected() int {
	return list.selected.Get()
}

func (list *List) SetSelected(selected int) {
	list.selected.Set(selected)
}

func (list *List) TemplateItem() Primitive {
	return list.templateItem.Get()
}

func (list *List) SetTemplateItem(item Primitive) {
	list.templateItem.Set(item)
}
