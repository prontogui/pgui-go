// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"errors"

	"github.com/prontogui/golib/pgcomm"
	"github.com/prontogui/golib/primitive"
)

type ProntoGUI interface {
	StartServing(addr string, port int) error
	StopServing()
	SetGUI(primitives ...primitive.Interface)
	Wait() (updatedPrimitive primitive.Interface, waitError error)
}

type _ProntoGUI struct {
	pgcomm     *pgcomm.PGComm
	synchro    *Synchro
	isgui      bool
	fullupdate bool
}

func (pg *_ProntoGUI) StartServing(addr string, port int) error {
	pg.fullupdate = true
	return pg.pgcomm.StartServing(addr, port)
}

func (pg *_ProntoGUI) StopServing() {
	pg.pgcomm.StopServing()
}

func (pg *_ProntoGUI) SetGUI(primitives ...primitive.Interface) {
	pg.fullupdate = true
	pg.isgui = true
	pg.synchro.SetTopPrimitives(primitives...)
}

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

func NewProntoGUI() ProntoGUI {
	pg := &_ProntoGUI{}

	pg.pgcomm = pgcomm.NewPGComm()
	pg.synchro = NewSynchro()
	pg.isgui = false
	pg.fullupdate = true

	return pg
}
