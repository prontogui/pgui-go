package field

type String struct {
	fieldno int8
	onset   func(int8)
	s       string
}

func (f *String) Get() string {
	return f.s
}

func (f *String) Set(s string) {
	f.s = s
	if f.onset != nil {
		f.onset(f.fieldno)
	}
}

func (f *String) OnSet(fieldno int8, onset func(int8)) {
	f.fieldno = fieldno
	f.onset = onset
}
