package field

type String struct {
	s string
}

func (f *String) Get() string {
	return f.s
}

func (f *String) Set(s string) {
	f.s = s
}
