package key

const INVALID_INDEX = -1

type PKey []int

func EmptyPKey() PKey {
	return PKey{}
}

func NewPKey(indices ...int) PKey {
	pk := make([]int, len(indices))
	copy(pk, indices)
	return pk
}

func NewPKeyFromAny(indices ...any) PKey {
	pk := make([]int, len(indices))
	for level, index := range indices {
		pk[level] = int(index.(uint64))
	}
	return pk
}

func (pk PKey) EqualTo(topk PKey) bool {
	if len(pk) != len(topk) {
		return false
	}
	for i, v := range pk {
		if v != topk[i] {
			return false
		}
	}
	return true
}

func (pk PKey) AddLevel(index int) PKey {
	origlen := len(pk)
	newpk := make([]int, origlen+1)
	copy(newpk, pk)
	newpk[origlen] = index
	return newpk
}

func (pk PKey) DescendsFrom(frompkey PKey) bool {
	if len(pk) <= len(frompkey) {
		return false
	}

	for level, index := range frompkey {
		if pk[level] != index {
			return false
		}
	}
	return true
}

func (pk PKey) IndexAtLevel(level int) int {
	if level < 0 || level >= len(pk) {
		return INVALID_INDEX
	}
	return pk[level]
}

type PKeyLocator struct {
	PKey          PKey
	LocationLevel int
}

func NewPKeyLocator(pkey PKey) *PKeyLocator {
	loc := &PKeyLocator{}
	loc.PKey = pkey
	loc.LocationLevel = -1
	return loc
}

func (loc *PKeyLocator) NextIndex() int {

	if loc.LocationLevel >= (len(loc.PKey) - 1) {
		panic("PKeyLocator went out of bounds")
	}

	loc.LocationLevel = loc.LocationLevel + 1

	return loc.PKey[loc.LocationLevel]
}

func (loc *PKeyLocator) Located() bool {
	return loc.LocationLevel == (len(loc.PKey) - 1)
}
