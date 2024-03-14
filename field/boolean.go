package field

type Boolean struct {
	b bool
}

func (f *Boolean) Get() bool {
	return f.b
}

func (f *Boolean) Set(b bool) {
	f.b = b
}
