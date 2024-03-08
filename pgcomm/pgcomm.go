package pgcomm

import (
	"fmt"
	"io"
	"log/slog"
	"net"

	pb "github.com/prontogui/golib/pb"
	"google.golang.org/grpc"
)

// The active server.
// - a valid reference after calling Serve() and returns no error.
// - null reference after calling StopServing().
var active_server *grpc.Server

// Implementation of the PGServer
type PGServerImpl struct {
	pb.UnimplementedPGServiceServer
}

var inboundUpdates chan []byte
var outboundUpdates chan []byte

// Package initializer
func init() {
	inboundUpdates = make(chan []byte, 2)
	outboundUpdates = make(chan []byte, 2)
}

// Called as a GO routine, this function pulls updates from a channel and streams them to app.
func streamOutboundUpdates(cancel chan bool, stream pb.PGService_StreamUpdatesServer) {

	// Loop for every udpate streamed back to caller...
	for {
		select {
		// If API call StreamUpdates cancels
		case <-cancel:
			return

		// When a new update is queued up or channel is closed...
		case cbor, ok := <-outboundUpdates:
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
func ExchangeUpdates(cborOut []byte) (bool, []byte) {

	// Queue the update to be streamed to app
	outboundUpdates <- cborOut

	// Wait for an update from app
	select {
	case cborIn, ok := <-inboundUpdates:
		if !ok {
			return false, nil
		}

		return true, cborIn
	}
}

// Implementation of PGServer.StreamUpdates API call
func (s *PGServerImpl) StreamUpdates(stream pb.PGService_StreamUpdatesServer) error {

	// Launch a Go routine to stream mutations back to caller
	cancel := make(chan bool)
	go streamOutboundUpdates(cancel, stream)

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

		inboundUpdates <- uxs.Cbor
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
func StartServing(addr string, port int) error {

	address := fmt.Sprintf("%s:%d", addr, port)

	lis, err := net.Listen("tcp", address)

	if err != nil {
		slog.Error("could not listen for network connection", "address", address, "error", err)
		return err
	}

	active_server := grpc.NewServer()

	pb.RegisterPGServiceServer(active_server, &PGServerImpl{})

	slog.Info("server is now listening", "address", address)

	go func() {
		if err := active_server.Serve(lis); err != nil {
			slog.Error("error occurred while serving", "address", address, "error", err)
		}
	}()

	return nil
}

// Stops serving of gRPC calls.
func StopServing() {
	if active_server != nil {
		active_server.GracefulStop()
	}
}

// Function to show a new GUI and wait for an update.

// Function to update the GUI with any changes and wait for the next update.
