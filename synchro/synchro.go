package synchro

import (
	"github.com/prontogui/golib/primitive"
)

type Synchro struct {
}

func NewSynchro() *Synchro {
	return &Synchro{}
}

func (s *Synchro) SetTopPrimitive(primitive.Primitive) {

}

func (s *Synchro) GetPendingUpdate() interface{} {
	return nil
}

func (s *Synchro) GetFullUpdate() interface{} {
	//l := []primitive.Primitive{}
	//	c := &primitive.Command{}

	//	l := []interface{}{c}

	//	return []interface{}{true, l}

	return nil
}

func (s *Synchro) AbsorbUpdate(update interface{}) {

}
