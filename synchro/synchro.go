package synchro

type Synchro struct {
	primitives []any
}

func NewSynchro() *Synchro {
	return &Synchro{}
}

func (s *Synchro) SetTopPrimitives(primitives []any) {
	s.primitives = primitives
}

func (s *Synchro) GetPendingUpdate() any {
	return nil
}

func (s *Synchro) GetFullUpdate() any {

	return []any{true, s.primitives}
}

func (s *Synchro) AbsorbUpdate(update any) {

}
