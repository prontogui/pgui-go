package golib

import (
	"errors"
	"fmt"

	cbor "github.com/fxamacker/cbor/v2"
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
		if update.pkey.EqualTo(pkey) && !update.ignored {
			return update
		}
	}
	return nil
}

func ignoreDescendantUpdates(updates []*Update, pkey key.PKey) {
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

	locator := key.NewPKeyLocator(pkey)

	// Get one of the top-level primitives to start with
	next := primitives[locator.NextIndex()]

	for !locator.Located() {
		// Try finding a descendant at the next level down
		next = next.LocateNextDescendant(locator)
	}

	return next
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
		ignoreDescendantUpdates(s.pendingUpdates, pkey)
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

func (s *Synchro) GetTopPrimitives() []primitive.Interface {
	return s.primitives
}

func (s *Synchro) GetPartialUpdate() ([]byte, error) {

	if len(s.pendingUpdates) == 0 {
		return nil, nil
	}

	updateList := []any{false}

	for _, update := range s.pendingUpdates {
		if !update.ignored {

			// Locate the primitive
			found := locatePrimitive(s.primitives, update.pkey)

			m := found.EgestUpdate(false, update.fields)

			// Add pkey and map to array of updates
			updateList = append(updateList, update.pkey, m)
		}
	}

	// Clear the pending updates
	s.pendingUpdates = []*Update{}

	return cbor.Marshal(updateList)
}

func (s *Synchro) GetFullUpdate() ([]byte, error) {

	if s.primitives == nil {
		return nil, nil
	}

	l := []any{true}

	for _, p := range s.primitives {
		p.EgestUpdate(true, nil)
		l = append(l, p.EgestUpdate(true, nil))
	}

	return cbor.Marshal(l)
}

func (s *Synchro) IngestUpdate(updatesCbor []byte) (updatedPrimitive primitive.Interface, updateError error) {

	var updates any

	updateError = cbor.Unmarshal(updatesCbor, &updates)
	if updateError != nil {
		return
	}

	var ok bool

	// Expecting a list of interfaces
	updatesList, ok := updates.([]any)
	if !ok {
		updateError = errors.New("the unmarshalled updates do not represent a list.  Expecting a list of updates")
		return
	}

	numitems := len(updatesList)

	// Must have length >= 1
	if len(updatesList) == 0 {
		updateError = errors.New("update must have atleast one value, the full/partial update flag")
		return
	}

	// Parse the full/partial update flag
	isfull, ok := updatesList[0].(bool)
	if !ok {
		updateError = errors.New("update value for full/partial flag is incorrect.  Expecting a bool")
		return
	}

	if isfull {
		updateError = errors.New("ingestion of full updates is not supported")
		return
	}

	// It's okay to have an empty partial update
	if numitems == 1 {
		return
	}

	if numitems != 3 {
		updateError = errors.New("partial update is limited to one primitive")
		return
	}

	// Get the pkey
	pkeyany, ok := updatesList[1].([]any)
	if !ok {
		updateError = errors.New("unable to convert pkey item to PKey")
		return
	}

	// Get the update map
	m, ok := updatesList[2].(map[any]any)
	if !ok {
		updateError = errors.New("unable to convert update item to map[any]any")
		return
	}

	pkey := key.NewPKeyFromAny(pkeyany...)
	updatedPrimitive = locatePrimitive(s.primitives, pkey)
	if updatedPrimitive == nil {
		updateError = fmt.Errorf("primitive at pkey = %v was not found", pkey)
		return
	}

	updateError = updatedPrimitive.IngestUpdate(m)

	return

}
