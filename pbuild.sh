#!/bin/sh

protoc --proto_path=proto --go-grpc_out=./ proto/clients.proto
protoc --proto_path=proto --go_out=./ proto/clients.proto
protoc-go-inject-tag -input=internal/api/clients.pb.go