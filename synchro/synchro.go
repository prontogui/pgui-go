package synchro

import (
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/prontogui/golib/primitive"
)

type Synchro struct {
	primitives []primitive.Primitive
}

func NewSynchro() *Synchro {
	return &Synchro{}
}

func (s *Synchro) SetTopPrimitives(primitives ...primitive.Primitive) {
	s.primitives = primitives

	// TODO:  assign pkeys to each primitive along with an onset closure

	// TODO:  add a means to track which primitives have been updated and track an "update number"

	/*
		for a, b := range primitives {

		}
	*/
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
