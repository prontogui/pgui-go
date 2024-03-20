package field

// TODO:  swap any type with primitive.Interface and update the test accordingly.
type Any2D struct {
	Reserved
	aaa [][]any
}

func (f *Any2D) Get() [][]any {
	return f.aaa
}

func (f *Any2D) Set(aaa [][]any) {
	f.aaa = aaa
	f.OnSet(true)
}
