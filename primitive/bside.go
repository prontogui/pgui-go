package primitive

import (
	"github.com/prontogui/golib/field"
)

type BSide struct {
	// The row index (0-based) where the primitive resides in a container (usually a table).
	// This is assigned automatically by the container primitive, or in the case of a
	// template row, by a mutation from the App after an event occured.
	Row field.Integer

	// The colunn index (0-based) where the primitive resides in a container (usually
	// a table).  This is assigned automatically by the container primitive, or in the case
	// of a template row, by a mutation from the App after an event occured.
	Col field.Integer

	// A JSON string specifying the embodiment to use for this primitive.
	// This can be assigned explicitly or by a mutation from the App.
	Embodiment field.String
}
