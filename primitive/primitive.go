package primitive

type Primitive interface {
	GetPKey() []uint
	AssignPKey([]uint, func([]uint) uint8)
}
