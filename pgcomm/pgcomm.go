// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pgcomm

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"

	pb "github.com/prontogui/golib/pb"
	"google.golang.org/grpc"
)

// Implementation of the PGServer
type PGComm struct {
	pb.UnimplementedPGServiceServer

	// The active server.
	// - a valid reference after calling Serve() and returns no error.
	// - null reference after calling StopServing().
	active_server *grpc.Server

	inboundUpdates  chan []byte
	outboundUpdates chan []byte
}

func NewPGComm() *PGComm {
	pgc := &PGComm{}

	pgc.inboundUpdates = make(chan []byte, 2)
	pgc.outboundUpdates = make(chan []byte, 2)

	return pgc
}

// Called as a GO routine, this function pulls updates from a channel and streams them to app.
func (pgc *PGComm) streamOutboundUpdates(cancel chan bool, stream pb.PGService_StreamUpdatesServer) {

	// Loop for every udpate streamed back to caller...
	for {
		select {
		// If API call StreamUpdates cancels
		case <-cancel:
			return

		// When a new update is queued up or channel is closed...
		case cbor, ok := <-pgc.outboundUpdates:
			if !ok {
				return
			}

			// Package it up into a protobuf structure and stream it out
			uxs := &pb.PGUpdate{Cbor: cbor}
			stream.SendMsg(uxs)
		}
	}
}

// Sends an update (as CBOR) to the app and waits for an update to come back.
func (pgc *PGComm) ExchangeUpdates(updateOut []byte) ([]byte, error) {

	// Queue the update to be streamed to app
	pgc.outboundUpdates <- updateOut

	// Wait for an update from app
	updateIn, ok := <-pgc.inboundUpdates
	if !ok {
		return nil, errors.New("inboundUpdates channel is invalid")
	}

	/*
		var cborIn interface{}

		err = cbor.Unmarshal(updateIn, &cborIn)
		if err != nil {
			return nil, err
		}
	*/

	return updateIn, nil
}

// Implementation of PGServer.StreamUpdates API call.
// TODO:  properly handle simultaneous calls to this API, e.g. two different apps connected
// at same time.
func (pgc *PGComm) StreamUpdates(stream pb.PGService_StreamUpdatesServer) error {

	// Launch a Go routine to stream mutations back to caller
	cancel := make(chan bool)
	go pgc.streamOutboundUpdates(cancel, stream)

	// Receive mutations and process them
	var err error

	// Loop for every inbound update received...
	for {
		uxs := pb.PGUpdate{}
		err = stream.RecvMsg(&uxs)

		if err != nil {
			break
		}

		// TODO:  This is mainly a debugging aid.  It should be removed for production.
		fmt.Printf("Update received with %d bytes.\n", len(uxs.Cbor))

		pgc.inboundUpdates <- uxs.Cbor
	}

	// End the GO routine for streaming outbound updates
	cancel <- true

	if err == io.EOF {
		return nil
	}

	return err
}

// Starts serving for gRPC calls at specified address and port.  Returns an error if it has
// problems opening a port for listening.
func (pgc *PGComm) StartServing(addr string, port int) error {

	address := fmt.Sprintf("%s:%d", addr, port)

	lis, err := net.Listen("tcp", address)

	if err != nil {
		slog.Error("could not listen for network connection", "address", address, "error", err)
		return err
	}

	pgc.active_server = grpc.NewServer()

	pb.RegisterPGServiceServer(pgc.active_server, pgc)

	slog.Info("server is now listening", "address", address)

	go func() {
		if err := pgc.active_server.Serve(lis); err != nil {
			slog.Error("error occurred while serving", "address", address, "error", err)
		}
	}()

	return nil
}

// Stops serving of gRPC calls.
func (pgc *PGComm) StopServing() {
	if pgc.active_server != nil {
		pgc.active_server.GracefulStop()
	}
}
