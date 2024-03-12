# pgui-go
Go language module for Pronto!GUI

## Generating the protobuf code
In the project folder:
```
make
```

## Setting up the gRPC tools

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

```

Add to ~/.zshrc or ~/.zprofile
```
export PATH="$PATH:$(go env GOPATH)/bin"
```

Refer to gRPC website for latest instructions:  https://grpc.io/docs/languages/go/quickstart/

