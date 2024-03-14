package field

type Integer struct {
	i int
}

func (f *Integer) Get() int {
	return f.i
}

func (f *Integer) Set(i int) {
	f.i = i
}
