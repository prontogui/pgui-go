package primitive

type Strings1DField interface {
	Get() []string
	Set([]string)
}

type BlobField interface {
	Get() []byte
	Set([]byte)
}

type PrimitiveField interface {
	Get() *Primitive
	Set(*Primitive)
}

type Primitives1DField interface {
	Get() []*Primitive
	Set([]*Primitive)
}

type Primitives2DField interface {
	Get() [][]*Primitive
	Set([][]*Primitive)
}
