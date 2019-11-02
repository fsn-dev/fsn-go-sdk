#!/bin/bash

set -eu

if [ ! -f "scripts/build.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

if [ $# -lt 1 ]; then
    echo "Usage: $0 <project> [<project>...]"
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

ignored_dirs="bin efsn scripts vendor"
for project in "$@"; do
    project=$(echo $project | sed 's#/*$##')
    [[ " $ignored_dirs " =~ " $project " ]] && continue
    build_project $project
done

#/* vim: set ts=4 sts=4 sw=4 et : */
