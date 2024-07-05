// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

type GroupWith struct {
	Embodiment string
	GroupItems []Primitive
}

func (w GroupWith) Make() *Group {
	grp := &Group{}
	grp.embodiment.Set(w.Embodiment)
	grp.groupItems.Set(w.GroupItems)
	return grp
}

type Group struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	groupItems Any1DField
}

func (grp *Group) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	grp.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &grp.embodiment},
			{key.FKey_GroupItems, &grp.groupItems},
		}
	})
}

// TODO:  generalize this code by handling inside primitive Reserved area.
func (grp *Group) LocateNextDescendant(locator *key.PKeyLocator) Primitive {
	if locator.NextIndex() != 0 {
		panic("cannot locate descendent using a pkey that we assumed was valid")
	}
	return grp.GroupItems()[locator.NextIndex()]
}

func (grp *Group) Embodiment() string {
	return grp.embodiment.Get()
}

func (grp *Group) SetEmbodiment(s string) {
	grp.embodiment.Set(s)
}

func (grp *Group) GroupItems() []Primitive {
	return grp.groupItems.Get()
}

func (grp *Group) SetGroupItems(items []Primitive) {
	grp.groupItems.Set(items)
}

func (grp *Group) SetGroupItemsVA(items ...Primitive) {
	grp.groupItems.Set(items)
}
