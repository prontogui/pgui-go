package golib

import (
	"github.com/prontogui/golib/field"
)

type BSide struct {

	// The row index (0-based) where the primitive resides in a container (usually a table).
	// This is assigned automatically by the container primitive, or in the case of a
	// template row, by a mutation from the App after an event occured.
	row field.Integer

	// The colunn index (0-based) where the primitive resides in a container (usually
	// a table).  This is assigned automatically by the container primitive, or in the case
	// of a template row, by a mutation from the App after an event occured.
	col field.Integer

	// A JSON string specifying the embodiment to use for this primitive.
	// This can be assigned explicitly or by a mutation from the App.
	embodiment field.String
}

func (bs *BSide) AttachFields(res *Reserved) {
	/*
		res.AttachField("B.Col", &bs.col, PKeyIndexDontCare)
		res.AttachField("B.Embodiment", &bs.embodiment, PKeyIndexDontCare)
		res.AttachField("B.Row", &bs.row, PKeyIndexDontCare)
	*/
}

func (bs *BSide) Row() int {
	return bs.row.Get()
}

func (bs *BSide) SetRow(i int) {
	bs.row.Set(i)
}

func (bs *BSide) Col() int {
	return bs.col.Get()
}

func (bs *BSide) SetCol(i int) {
	bs.col.Set(i)
}

func (bs *BSide) Embodiment() string {
	return bs.embodiment.Get()
}

func (bs *BSide) SetEmbodiment(s string) {
	bs.embodiment.Set(s)
}
