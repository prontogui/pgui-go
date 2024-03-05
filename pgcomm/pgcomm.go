package pgcomm

import (
	"fmt"
	"log/slog"
	"net"

	pb "github.com/prontogui/golib/pb"
	"google.golang.org/grpc"
)

// The active server.
// - a valid reference after calling Serve() and returns no error.
// - null reference after calling StopServing().
var active_server *grpc.Server

type server struct {
	pb.UnimplementedPGServiceServer
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

	pb.RegisterPGServiceServer(active_server, &server{})

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
