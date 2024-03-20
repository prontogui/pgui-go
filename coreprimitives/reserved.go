package coreprimitives

const (
	// The maximum number of fields in any given primitive.  TODO:  check for accuracy of this in unit testing,
	// in case a primitive is updated or added without changing this number.
	MaxPrimitiveFields = 4
)

/*
Reserved fields for primitive updates.
*/
type Reserved struct {
}
