package field

type Any1D struct {
	aa []any
}

func (f *Any1D) Get() []any {
	return f.aa
}

func (f *Any1D) Set(aa []any) {
	f.aa = aa
}
