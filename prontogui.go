// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"errors"

	"github.com/prontogui/golib/pgcomm"
	"github.com/prontogui/golib/primitive"
)

// The API for working with app-to-server communication.
type ProntoGUI interface {
	StartServing(addr string, port int) error
	StopServing()
	SetGUI(primitives ...primitive.Interface)
	Wait() (updatedPrimitive primitive.Interface, waitError error)
}

// Internal data for handling the API of this library
type _ProntoGUI struct {
	pgcomm     *pgcomm.PGComm
	synchro    *Synchro
	isgui      bool
	fullupdate bool
}

// Start up the communication server so we can accept a connection from the app.
func (pg *_ProntoGUI) StartServing(addr string, port int) error {
	pg.fullupdate = true
	return pg.pgcomm.StartServing(addr, port)
}

// Stop the communication and close any open connection.
func (pg *_ProntoGUI) StopServing() {
	pg.pgcomm.StopServing()
}

// Prepares all the primitives that will participate in the GUI.  This must be called at least once before Wait().
//
// Note:  it is possible to reuse primitive instance throughout the same GUI and even across multiple GUIs but there are
// some things to be aware of.
//
//   - If you change a field of a primitive that no longer is part of the active GUI then you could have indesirable behavior
//     or a panic situation.  This is because the primitive is still "prepped" for communication but has since been divorced
//     from the GUI model.  Right now, there is no compelling reason to "unprepare" primitives that no longer participte in
//     the GUI since this would have some impact on performance.  This could change as ProntoGUI usage evolves and perhaps
//     we might add a flexible policy in the future.  For now, don't do anything with non-participating primitives until
//     they're part of the active GUI again (by calling SetGUI).
//
//   - Updates to the primitive coming from the app side will not be discernable.  To explain this, let's say you have the
//     same Command primitive instance in three different views.  When the user clicks on the Command, you will receive
//     on update from the Wait() function that the command was clicked.  However, you be able to know which view it was
//     participating in when the click happened.
//
//   - Updates to the primitive coming from the server side will only affect one instance of the primitive on the app side.
//     For example, if you change the label of a Command primitive that is reused several times in the same GUI, only one of
//     them will visually change.
//
//   - It's best to avoid reusing primitives in the same GUI unless they are static, such as Text.
func (pg *_ProntoGUI) SetGUI(primitives ...primitive.Interface) {
	pg.fullupdate = true
	pg.isgui = true
	pg.synchro.SetTopPrimitives(primitives...)
}

// Sends model updates to the app for rendering, waits for an event to occured in the app that requires server attention,
// and receives model updates from the app.  This function returns the primitive that ended triggered the event that ended
// the wait period.  An error is returend if something went wrong.
func (pg *_ProntoGUI) Wait() (updatedPrimitive primitive.Interface, waitError error) {

	if !pg.isgui {
		return nil, errors.New("no GUI has been set")
	}

	var updateOut []byte
	var updateIn []byte

	// Need to send a full update?
	if pg.fullupdate {
		updateOut, waitError = pg.synchro.GetFullUpdate()
		pg.fullupdate = false
	} else {
		updateOut, waitError = pg.synchro.GetPartialUpdate()
	}
	if waitError != nil {
		return
	}

	updateIn, waitError = pg.pgcomm.ExchangeUpdates(updateOut)
	if waitError != nil {
		return
	}

	return pg.synchro.IngestUpdate(updateIn)
}

// Creates a new ProntoGUI instance.
func NewProntoGUI() ProntoGUI {
	pg := &_ProntoGUI{}

	pg.pgcomm = pgcomm.NewPGComm()
	pg.synchro = NewSynchro()
	pg.isgui = false
	pg.fullupdate = true

	return pg
}
