package key

type PKey uint64

func (pk *PKey) AddLevel(index int) PKey {
	// TODO:  implement this
	return PKey(index)
}

func (pk *PKey) DescendsFrom(pkey PKey) bool {
	// TODO:  implement this
	return false
}

func (pk *PKey) IndexAtLevel(level int) int {

	if level == 0 {
		return int(*pk)
	}
	return -1
}
