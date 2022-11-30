#!/bin/bash
CURRENT_DIR=$PWD
rm -rf ${CURRENT_DIR}/pkg/proto/*
for x in $(find ${CURRENT_DIR}/protos/*); do
  protoc -I=${CURRENT_DIR}/protos --go_out=${CURRENT_DIR}/pkg/proto --go-grpc_out=${CURRENT_DIR}/pkg/proto ${x}
done
