
## 1. Install protoc-gen-go
`go get google.golang.org/protobuf/cmd/protoc-gen-go`
`go install google.golang.org/protobuf/cmd/protoc-gen-go`

## Update PATH enviroment
`export PATH=$PATH:$GOPATH/bin`

## Generate go code
`protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto`