package field

type Integer struct {
	Reserved
	i int
}

func (f *Integer) Get() int {
	return f.i
}

func (f *Integer) Set(i int) {
	f.i = i
	f.OnSet(false)
}
