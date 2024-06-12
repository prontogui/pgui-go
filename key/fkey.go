// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package key

type FKey uint8

const (
	INVALID_FIELDNAME = 255
	INVALID_FKEY      = ""
)

const (

	// ADD NEW FIELDS TO THIS BLOCK - ALPHABETICAL ORDER PLEASE!
	FKey_Changed FKey = iota
	FKey_Checked
	FKey_Choice
	FKey_Choices
	FKey_Content
	FKey_Data
	FKey_Embodiment
	FKey_FrameItems
	FKey_GroupItems
	FKey_Issued
	FKey_Label
	FKey_ListItems
	FKey_Rows
	FKey_Selected
	FKey_Showing
	FKey_State
	FKey_Status
	FKey_TemplateItem
	FKey_TextEntry

	// RESERVED CONSTANT
	FKey_MAXIMUMKEYS
)

var _fkeyToName []string
var _nameToFKey map[string]FKey

func init() {
	_fkeyToName = make([]string, FKey_MAXIMUMKEYS)

	// ADD NEW FIELDS TO THIS BLOCK - ALPHABETICAL ORDER PLEASE!
	_fkeyToName[FKey_Changed] = "Changed"
	_fkeyToName[FKey_Checked] = "Checked"
	_fkeyToName[FKey_Choice] = "Choice"
	_fkeyToName[FKey_Choices] = "Choices"
	_fkeyToName[FKey_Content] = "Content"
	_fkeyToName[FKey_Data] = "Data"
	_fkeyToName[FKey_Embodiment] = "Embodiment"
	_fkeyToName[FKey_FrameItems] = "FrameItems"
	_fkeyToName[FKey_GroupItems] = "GroupItems"
	_fkeyToName[FKey_Issued] = "Issued"
	_fkeyToName[FKey_Label] = "Label"
	_fkeyToName[FKey_ListItems] = "ListItems"
	_fkeyToName[FKey_Rows] = "Rows"
	_fkeyToName[FKey_Selected] = "Selected"
	_fkeyToName[FKey_Showing] = "Showing"
	_fkeyToName[FKey_State] = "State"
	_fkeyToName[FKey_Status] = "Status"
	_fkeyToName[FKey_TemplateItem] = "TemplateItem"
	_fkeyToName[FKey_TextEntry] = "TextEntry"

	_nameToFKey = make(map[string]FKey, FKey_MAXIMUMKEYS)

	for fkey, fname := range _fkeyToName {
		_nameToFKey[fname] = FKey(fkey)
	}
}

func FKeyFor(fieldname string) FKey {

	fkey, ok := _nameToFKey[fieldname]

	if ok {
		return fkey
	}

	return INVALID_FIELDNAME
}

func FieldnameFor(fkey FKey) string {

	if fkey >= FKey_MAXIMUMKEYS {
		return INVALID_FKEY
	}

	return _fkeyToName[fkey]
}
