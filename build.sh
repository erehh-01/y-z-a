#!/bin/bash

ROOT_PATH=$(pwd)
BUILD_PATH=$ROOT_PATH/build/
LIB_PATH=$ROOT_PATH/lib/runtime/linux/
APP_NAME=Y-Z-A

rm -rf $BUILD_PATH

mkdir -p "$BUILD_PATH"

cp -r "$LIB_PATH" "$BUILD_PATH"

export GOARCH=amd64
export APP_NAME=Y-Z-A-64
export CGO_CFLAGS="-I$ROOT_PATH/lib/build/linux/include"
export LD_LIBRARY_PATH="-L$ROOT_PATH/lib/build/linux/lib:$LD_LIBRARY_PATH"
export CGO_LDFLAGS="-L$ROOT_PATH/lib/build/linux/lib  -ltdjson_static -ltdjson_private -ltdclient -ltdcore -ltdactor -ltdapi -ltddb -ltdsqlite -ltdnet -ltdutils -lssl -lcrypto"
export CGO_ENABLED=1
export GOOS=linux

go build -v \
    -ldflags "-w -s" \
    -tags netgo \
    -gcflags="-S -m" \
    -trimpath -mod=readonly \
    -buildmode=pie -a -installsuffix cgo \
    -o "$BUILD_PATH/linux/$APP_NAME" .

cd "$ROOT_PATH"
