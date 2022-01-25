// Package proto holds the protobuf and corresponding generated code for the broker service
//go:generate protoc -I . --go_out=. --go-grpc_out=. broker.proto
package proto
