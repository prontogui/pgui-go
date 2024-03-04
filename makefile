# Makefile to generate all code for .proto files.

# Define the protobuf compiler and the grpc plugin
PROTOC = protoc

all: 
	git clone -b v0.0.2 https://andyhjoseph@github.com/prontogui/proto.git
	mkdir -p pb
	$(PROTOC) --go_opt=paths=import --go_out=pb proto/pg.proto 
	$(PROTOC) --go-grpc_out=pb proto/pg.proto

.PHONY: all

clean:
	rm -f *.go
