#!/bin/bash

set -eu

if [ ! -f "scripts/build.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

if (( $# < 1 )) || [[ "$1" = "-h" ]] || [[ "$1" = "--help" ]] ; then
    echo "Usage: $0 <project> [<project>...]"
    exit
fi

if ! command -v go > /dev/null; then
    echo "please install go"
    exit
fi

version_gt() {
    test "$(printf '%s\n' "$@" | sort -V | head -n 1)" != "$1"
}
golang_version=$(go version |cut -d' ' -f3 |sed 's/go//')
if (version_gt 1.11 $golang_version); then
    echo "go version should be greater than or equal to 1.11"
    exit
fi

build_project() {
    echo "----------- build $project -----------"
    main_file="./$project/main.go"
    if [ ! -f "$main_file" ]; then
        echo "Error: $main_file not exist"
        return
    fi
    echo "RUN go build -v -mod=vendor -o bin/$project $project/*.go"
    go build -v -mod=vendor -o bin/$project $project/*.go
    echo "Build finished, run \"./bin/$project\" to launch."
}

if [[ ! -d vendor ]]; then
    echo "please 'make vendor' or 'make vendor_with_proxy' firstly"
    exit
fi

ignored_dirs="bin efsn scripts vendor fsnapi"
for project in "$@"; do
    project=$(echo $project | sed 's#/*$##')
    [[ " $ignored_dirs " =~ " $project " ]] && continue
    build_project $project
done

#/* vim: set ts=4 sts=4 sw=4 et : */
