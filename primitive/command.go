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
