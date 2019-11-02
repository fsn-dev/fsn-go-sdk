#!/bin/bash

set -eu

if [ ! -f "scripts/gomod.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

export GO111MODULE=on
export GOPROXY=https://goproxy.io

echo "RUN go mod init"
go mod init github.com/FusionFoundation/fsn-go-sdk 2>/dev/null || true

echo "Run go mod vendor"
go mod vendor -v

#/* vim: set ts=4 sts=4 sw=4 et : */
