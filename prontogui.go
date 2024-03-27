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
	Wait() error
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
	pg.synchro.SetTopPrimitives(primitives...)
}

func (pg *_ProntoGUI) Wait() error {

	if !pg.isgui {
		return errors.New("no GUI has been set")
	}

	var updateOut []byte
	var err error
	var updateIn []byte

	// Need to send a full update?
	if pg.fullupdate {

		updateOut, err = NewSynchro().GetFullUpdate()
		if err != nil {
			return err
		}

		pg.fullupdate = false
	}

	updateIn, err = pg.pgcomm.ExchangeUpdates(updateOut)
	if err != nil {
		return err
	}

	NewSynchro().IngestUpdate(updateIn)

	return nil
}

func NewProntoGUI() ProntoGUI {
	pg := &_ProntoGUI{}

	pg.pgcomm = pgcomm.NewPGComm()
	pg.synchro = NewSynchro()
	pg.isgui = false
	pg.fullupdate = true

	return pg
}
