package field

type Boolean struct {
	Reserved
	b bool
}

func (f *Boolean) Get() bool {
	return f.b
}

func (f *Boolean) Set(b bool) {
	f.b = b
	f.OnSet()
}
