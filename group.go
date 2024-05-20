package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

type GroupWith struct {
	GroupItems []primitive.Interface
}

func (w GroupWith) Make() *Group {
	grp := &Group{}
	grp.groupItems.Set(w.GroupItems)
	return grp
}

type Group struct {
	// Mix-in the common guts for primitives (B-side fields, implements primitive interface, etc.)
	Reserved

	groupItems field.Any1D
}

func (grp *Group) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	grp.AttachField("GroupItems", &grp.groupItems)

	// Prepare all fields for updates
	grp.Reserved.PrepareForUpdates(pkey, onset)
}

func (grp *Group) LocateNextDescendant(locator *key.PKeyLocator) primitive.Interface {
	return grp.GroupItems()[locator.NextIndex()]
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
