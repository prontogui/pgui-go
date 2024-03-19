package synchro

import (
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

type Synchro struct {
	primitives []primitive.Primitive

	pendingUpdates map[key.PKey][primitive.MaxPrimitiveFields]key.FKey
}

func NewSynchro() *Synchro {
	return &Synchro{}
}

func (s *Synchro) OnSet(pkey key.PKey, fkey key.FKey) {

	fkeys := s.pendingUpdates[pkey]

	var nextfree int
	var nextfkey key.FKey

	for nextfree, nextfkey = range fkeys {
		// Already been recorded as an update?
		if nextfkey == fkey {
			return
		}
		// Found a free location to store fkey?
		if fkey == 0 {
			break
		}
	}

	// TODO:  test for case where every field is updated for the primitive with the most fields.  Make sure we
	// don't get index out of bounds.
	fkeys[nextfree] = fkey

	s.pendingUpdates[pkey] = fkeys
}

func (s *Synchro) SetTopPrimitives(primitives ...primitive.Primitive) {
	//	s.pendingUpdates = make(map[key.PKey][primitive.MaxPrimitiveFields]key.FKey)

	s.primitives = primitives

	var pkey key.PKey

	for i, p := range primitives {
		p.NotifyOnSet(pkey.AddLevel(i), s.OnSet)
	}
}

func (s *Synchro) GetPartialUpdate() []byte {
	return nil
}

func (s *Synchro) GetFullUpdate() []byte {

	cmd := primitive.Command{}
	cmd.Label.Set("Ha!")
	l := []any{true, &cmd}
	cbor, _ := cbor.Marshal(l)

	return cbor
}

func (s *Synchro) AbsorbUpdate(cbor []byte) {

}
