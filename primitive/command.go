package primitive

import (
	"github.com/prontogui/golib/field"
)

type Command struct {
	BSide  BSide
	Label  field.String
	Issued field.Boolean
	Status field.Integer
}

/*
type Command2 interface {
	X int;
	LabelGet() string;
	LabelSet(string);

	//Issued bool
	//Status int
	//Index  int
}
*/

// []byte for file content and images  --> BlobField
// []Primitive for containers  --> Primitives1DField
// [][]Primitive for tables --> Primitives2DField
// Single Primitive for template  --> PrimitiveField
// []string for choices  --> Strings1DField
// BSide

func SomeFunc() {
	c := Command{}

	c.Label.Set("asdf")
	c.Status.Set(0)
	c.BSide.Col.Get()
}
