package field

type Strings1D struct {
	Reserved
	sa []string
}

func (f *Strings1D) Get() []string {
	return f.sa
}

func (f *Strings1D) Set(sa []string) {
	f.sa = sa
	f.OnSet(false)
}
