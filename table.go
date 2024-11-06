// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"errors"
	"slices"

	"github.com/prontogui/golib/key"
)

// A table displays an array of primitives in a grid of rows and columns.
type TableWith struct {
	Embodiment string
	Headings   []string
	Rows       [][]Primitive
	Status     int
	Tag        string
}

// Creates a new Table using the supplied field assignments.
func (w TableWith) Make() *Table {
	table := &Table{}
	table.embodiment.Set(CanonizeEmbodiment(w.Embodiment))
	table.headings.Set(w.Headings)
	table.rows.Set(w.Rows)
	table.status.Set(w.Status)
	table.tag.Set(w.Tag)
	return table
}

// A table displays an array of primitives in a grid of rows and columns.
type Table struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	headings   Strings1DField
	rows       Any2DField
	status     IntegerField
	tag        StringField
}

// Creates a new Table with headings.
func NewTable(headings ...string) *Table {
	return TableWith{Headings: headings}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (table *Table) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	table.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &table.embodiment},
			{key.FKey_Headings, &table.headings},
			{key.FKey_Rows, &table.rows},
			{key.FKey_Status, &table.status},
			{key.FKey_Tag, &table.tag},
		}
	})
}

// A non-recursive method to locate descendants by PKey.  This is used internally by this library
// and normally should not be called by users of the library.
func (table *Table) LocateNextDescendant(locator *key.PKeyLocator) Primitive {
	// TODO:  generalize this code by handling inside primitive Reserved area.

	nextIndex := locator.NextIndex()

	// Fields are handled in alphabetical order
	switch nextIndex {
	case 0:
		// TODO:  Optimization - add a row/col accessor to Any2D field so we don't return all the contents just
		// to index a single item here.  Same could be done for Any1D.
		row := locator.NextIndex()
		col := locator.NextIndex()
		return table.Rows()[row][col]
	default:
		panic("cannot locate descendent using a pkey that we assumed was valid")
	}
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (table *Table) Embodiment() string {
	return table.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (table *Table) SetEmbodiment(s string) *Table {
	table.embodiment.Set(CanonizeEmbodiment(s))
	return table
}

// Returns the headings to use for each column in the table.
func (table *Table) Headings() []string {
	return table.headings.Get()
}

// Sets the headings to use for each column in the table.
func (table *Table) SetHeadings(s []string) *Table {
	table.headings.Set(s)
	return table
}

// Sets the headings (as variadic arguments) to use for each column in the table.
func (table *Table) SetHeadingsVA(items ...string) *Table {
	table.headings.Set(items)
	return table
}

// Returns the dynamically populated 2D (rows, cols) collection of primitives that appear in the table.
func (table *Table) Rows() [][]Primitive {
	return table.rows.Get()
}

// Sets the dynamically populated 2D (rows, cols) collection of primitives that appear in the table.
func (table *Table) SetRows(items [][]Primitive) *Table {
	table.rows.Set(items)
	return table
}

// Returns the status of the table:  0 = Table Normal, 1 = Table Disabled, 2 = Table Hidden.
func (table *Table) Status() int {
	return table.status.Get()
}

// Sets the status of the table:  0 = Table Normal, 1 = Table Disabled, 2 = Table Hidden.
func (table *Table) SetStatus(status int) *Table {
	table.status.Set(status)
	return table
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Tables inside other containers.
func (table *Table) Tag() string {
	return table.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Tables inside other containers.
func (table *Table) SetTag(s string) *Table {
	table.tag.Set(s)
	return table
}

// Inserts a new row in this table before the index specified.  If index is -1 or extends beyond the number
// of rows in the table then row is appended at the end of the table.
// The row must match the dimension and cell types of the template row
func (table *Table) InsertRow(index int, row []Primitive) {

	originalRows := table.rows.Get()

	if index < 0 || index > len(originalRows) {
		table.rows.Set(append(originalRows, row))
		return
	}

	table.rows.Set(slices.Insert(originalRows, index, row))
}

// Deletes a row in this table at the given index.  An error is returned if the index is out of range.
func (table *Table) DeleteRow(index int) error {

	originalRows := table.rows.Get()

	if index < 0 || index >= len(originalRows) {
		return errors.New("index out of range")
	}

	table.rows.Set(slices.Delete(originalRows, index, index+1))

	return nil
}

// Convenience function that returns the number of rows.
func (table *Table) RowCount() int {
	return table.rows.Length()
}

func (table *Table) DeleteAllRows() {
	table.rows.Set([][]Primitive{})
}
