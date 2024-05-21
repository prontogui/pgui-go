package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

type GroupWith struct {
	Embodiment string
	GroupItems []primitive.Interface
}

func (w GroupWith) Make() *Group {
	grp := &Group{}
	grp.embodiment.Set(w.Embodiment)
	grp.groupItems.Set(w.GroupItems)
	return grp
}

type Group struct {
	// Mix-in the common guts for primitives
	Reserved

	embodiment field.String
	groupItems field.Any1D
}

func (grp *Group) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	grp.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &grp.embodiment},
			{key.FKey_GroupItems, &grp.groupItems},
		}
	})
}

func (grp *Group) LocateNextDescendant(locator *key.PKeyLocator) primitive.Interface {
	return grp.GroupItems()[locator.NextIndex()]
}

func (grp *Group) Embodiment() string {
	return grp.embodiment.Get()
}

func (grp *Group) SetEmbodiment(s string) {
	grp.embodiment.Set(s)
}

func (grp *Group) GroupItems() []primitive.Interface {
	return grp.groupItems.Get()
}

func (grp *Group) SetGroupItems(items []primitive.Interface) {
	grp.groupItems.Set(items)
}

func (grp *Group) SetGroupItemsVA(items ...primitive.Interface) {
	grp.groupItems.Set(items)
}
