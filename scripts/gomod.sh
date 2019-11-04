#!/bin/bash

set -eu

if [ ! -f "scripts/gomod.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

print_uasge() {
    echo "Usage: $0 [--proxy]"
    echo ""
    echo "if '--proxy' is specified, then will set 'GOPROXY=https://goproxy.io'"
    echo "this is useful for someone who can not download packages from 'golang.org'"
}

useproxy=false

if (( $# > 0 )); then
    if [[ "$1" = "-h" || "$1" == "--help" ]]; then
        print_uasge
        exit
    fi

    if [[ "$1" = "--proxy" ]]; then
        useproxy=true
    fi
fi

export GO111MODULE=on
if [[ "$useproxy" = "true" ]]; then
    export GOPROXY=https://goproxy.io
fi

echo "RUN go mod init"
go mod init github.com/FusionFoundation/fsn-go-sdk 2>/dev/null || true

echo "Run go mod vendor"
go mod vendor -v

#/* vim: set ts=4 sts=4 sw=4 et : */
