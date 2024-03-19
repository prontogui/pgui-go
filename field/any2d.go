package field

type Any2D struct {
	Reserved
	aaa [][]any
}

func (f *Any2D) Get() [][]any {
	return f.aaa
}

func (f *Any2D) Set(aaa [][]any) {
	f.aaa = aaa
	f.OnSet()
}
