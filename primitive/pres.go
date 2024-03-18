package primitive

/*
Reserved fields for primitive updates.
*/
type PReserved struct {
	pkey         []uint // `cbor:"omitempty"`
	updateno     uint8
	onset        func([]uint) uint8
	fieldupdates []uint8
}

func (pr *PReserved) Init(numfields uint8) {
	pr.fieldupdates = make([]uint8, numfields+NumBsideFields)
}

func (pr *PReserved) GetPKey() []uint {
	return pr.pkey
}

func (pr *PReserved) AssignPKey(pkey []uint, onset func([]uint) uint8) {

	pr.pkey = pkey
	pr.onset = onset
}

func (pr *PReserved) OnFieldSet(fieldno int8) {

	if pr.onset != nil {
		// Relay up the chain, that a field was changed at specified pkey.
		pr.updateno = pr.onset(pr.pkey)
	}

	// Record that field was change at such and such update number
	pr.fieldupdates[fieldno+NumBsideFields] = pr.updateno
}
