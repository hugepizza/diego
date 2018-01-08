//go:generate sh -c "protoc -I/usr/local/include -I. -I$GOPATH/src --grpc-gateway_out=logtostderr=true:. --go_out=plugins=grpc:. *.proto"
// protoc --go_out=. *.proto
package protos
