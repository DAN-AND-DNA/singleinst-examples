#!/bin/sh
root=$(cd `dirname $0`; pwd)

#protoc  --proto_path=$root --go_out=../../..  --go-grpc_out=../../.. pb/proto/*.proto

go run $root/kvservice/main.go --config $root/kvservice/config/modules
