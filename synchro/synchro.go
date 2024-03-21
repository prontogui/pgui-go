package synchro

import (
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/prontogui/golib/coreprimitives"
	"github.com/prontogui/golib/key"
	"github.com/prontogui/golib/primitive"
)

type Update struct {
	pkey    key.PKey
	fields  []key.FKey
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

	for _, nextfkey := range update.fields {
		// Already been recorded as an update?
		if nextfkey == fkey {
			return
		}
	}

	update.fields = append(update.fields, fkey)
}

func locatePrimitive(primitives []primitive.Interface, pkey key.PKey) primitive.Interface {

	level := 0

	var found primitive.Interface

	// Get one of the top-level primitives to start with
	next := primitives[pkey.IndexAtLevel(level)]

	for next != nil {
		found = next

		// Try finding a child at the next level down
		level = level + 1
		next = found.GetChildPrimitive(pkey.IndexAtLevel(level))
	}

	return found
}

func (s *Synchro) OnSet(pkey key.PKey, fkey key.FKey, structural bool) {

	// is there pending update for this primitive?
	existingUpdate := findPendingUpdate(s.pendingUpdates, pkey)
	if existingUpdate != nil {
		appendFieldToUpdate(existingUpdate, fkey)
	} else {
		// Add a new update to pending
		newUpdate := &Update{pkey: pkey}
		newUpdate.fields = []key.FKey{fkey}
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

func marshalFieldsToMap(p primitive.Interface, fields []key.FKey) map[string]any {

	m := make(map[string]any, len(fields))

	for _, fkey := range fields {
		fieldname := key.FieldnameFor(fkey)
		m[fieldname] = p.GetFieldValue(fieldname)
	}

	return m
}

func (s *Synchro) GetPartialUpdate() ([]byte, error) {

	if len(s.pendingUpdates) == 0 {
		return nil, nil
	}

	var updateList []any

	for _, update := range s.pendingUpdates {
		if !update.ignored {

			// Locate the primitive
			found := locatePrimitive(s.primitives, update.pkey)

			m := marshalFieldsToMap(found, update.fields)

			// Add pkey and map to array of updates
			updateList = append(updateList, update.pkey, m)
		}
	}

	// Clear the pending updates
	s.pendingUpdates = []*Update{}

	return cbor.Marshal(updateList)
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
