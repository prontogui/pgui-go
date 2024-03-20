package synchro

import (
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/prontogui/golib/coreprimitives"
	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

type Update struct {
	pkey    key.PKey
	fields  [coreprimitives.MaxPrimitiveFields]key.FKey
	ignored bool
}

type Synchro struct {
	primitives []primitive.Interface

	pendingUpdates []*Update
}

func NewSynchro() *Synchro {
	return &Synchro{}
}

func findPendingUpdate(updates []*Update, pkey key.PKey) *Update {
	for _, update := range updates {
		if update.pkey == pkey && !update.ignored {
			return update
		}
	}
	return nil
}

func ignoreDescendentUpdates(updates []*Update, pkey key.PKey) {
	for _, update := range updates {
		if update.pkey.DescendsFrom(pkey) {
			update.ignored = true
		}
	}
}

func appendFieldToUpdate(update *Update, fkey key.FKey) {

	var nextfree int
	var nextfkey key.FKey

	for nextfree, nextfkey = range update.fields {
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
	update.fields[nextfree] = fkey
}

func (s *Synchro) OnSet(pkey key.PKey, fkey key.FKey, structural bool) {

	// is there pending update for this primitive?
	existingUpdate := findPendingUpdate(s.pendingUpdates, pkey)
	if existingUpdate != nil {
		appendFieldToUpdate(existingUpdate, fkey)
	} else {
		// Add a new update to pending
		newUpdate := &Update{pkey: pkey}
		newUpdate.fields[0] = fkey
		s.pendingUpdates = append(s.pendingUpdates, newUpdate)
	}

	if structural {
		ignoreDescendentUpdates(s.pendingUpdates, pkey)
	}
}

func (s *Synchro) SetTopPrimitives(primitives ...primitive.Interface) {
	//	s.pendingUpdates = make(map[key.PKey][primitive.MaxPrimitiveFields]key.FKey)

	s.primitives = primitives

	var pkey key.PKey

	for i, p := range primitives {
		p.PrepareForUpdates(pkey.AddLevel(i), s.OnSet)
	}
}

func (s *Synchro) GetPartialUpdate() []byte {

	if len(s.pendingUpdates) == 0 {
		return nil
	}

	for _, update := range s.pendingUpdates {
		if !update.ignored {
			// Locate the primitive

			// Get updates as a map from primitive based on fields

			// Add pkey and map to array of updates

		}
	}

	// Clear the pending updates
	s.pendingUpdates = []*Update{}

	return nil
}

func (s *Synchro) GetFullUpdate() []byte {

	cmd := coreprimitives.Command{}
	cmd.Label.Set("Ha!")
	l := []any{true, &cmd}
	cbor, _ := cbor.Marshal(l)

	return cbor
}

func (s *Synchro) IngestUpdates(cbor []byte) {

}
