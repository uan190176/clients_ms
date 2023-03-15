#!/bin/sh

protoc --proto_path=proto --go-grpc_out=./ proto/clients.proto
protoc --proto_path=proto --go_out=./ proto/clients.proto
protoc-go-inject-tag -input=internal/api/clients.pb.go

cp internal/api/clients.pb.go /Volumes/SSD/_goland/NormalizeAddress_MS/internal/api/clients.pb.go
cp internal/api/clients_grpc.pb.go /Volumes/SSD/_goland/NormalizeAddress_MS/internal/api/clients_grpc.pb.go
cp proto/clients.proto /Volumes/SSD/_goland/NormalizeAddress_MS/proto/clients.proto
