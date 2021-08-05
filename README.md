# golang-grpc


### Install the protocol compiler plugins for Go using the following commands:
###### Protocol buffer compiler, protoc, version 3.

```sh
go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get github.com/golang/protobuf@latest
```
###### Update your PATH so that the protoc compiler can find the plugins:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

Regenerate Proto-buffers in GO

protoc -I="." --go_out="./chat" chat.proto