package field

type Any2D struct {
	aaa [][]any
}

func (f *Any2D) Get() [][]any {
	return f.aaa
}

func (f *Any2D) Set(aaa [][]any) {
	f.aaa = aaa
}
