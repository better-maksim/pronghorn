#!/usr/bin/env bash

TARGET_FILE_NAME=pronghorn
SOURCE_FILE_NAME=main
build() {
  echo $GOOS $GOARCH
  tname=${TARGET_FILE_NAME}_${GOOS}_${GOARCH}${EXT}
  env GOOS=$GOOS GOARCH=${GOARCH} \
  go build -o ${tname} \
  -v ${SOURCE_FILE_NAME}.go
  chmod +x ${tname}
}
#mac os
GOOS=darwin
GOARCH=amd64
build

#linux
GOOS=linux
GOARCH=amd64
build

#windows
GOOS=windows
GOARCH=amd64
build

GOARCH=386
build