#!/bin/sh
root=$(cd `dirname $0`; pwd)

#protoc  --proto_path=$root --go_out=../../..  --go-grpc_out=../../.. pb/proto/*.proto

go run "$root"/webbff/main.go --config "$root"/webbff/config/modules
