# Copyright 2024 ProntoGUI, LLC.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# Makefile to generate all code for .proto files.

# Define the protobuf compiler and the grpc plugin
PROTOC = protoc

all: 
	if [ ! -d "proto" ]; then git clone -b v0.0.2 https://andyhjoseph@github.com/prontogui/proto.git; fi
	mkdir -p pb
	$(PROTOC) --go_opt=paths=import --go_out=pb proto/pg.proto 
	$(PROTOC) --go-grpc_out=pb proto/pg.proto

.PHONY: all

clean:
	rm -Rf proto
	rm -Rf pb
