#!/bin/sh
OUT_DIR=../../pkg/protobuf/
mkdir -p $OUT_DIR
OPTION="paths=source_relative"
protoc -I=./ \
    --go_out=$OPTION:$OUT_DIR \
    --go_opt=$OPTION \
    --go-grpc_out=require_unimplemented_servers=false,$OPTION:$OUT_DIR \
    --go-grpc_opt=require_unimplemented_servers=false,$OPTION \
    ./**/*.proto;
