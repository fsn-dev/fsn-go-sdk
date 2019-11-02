#!/bin/bash

set -eu

if [ ! -f "scripts/run.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

if (( $# < 1 )) || [[ $1 =~ ^- ]]; then
    echo "Usage: $0 <project> [options]"
    exit
fi

project=$(echo $1 | sed 's#/*$##')
shift

bin_file="./bin/$project"
if [ ! -f "$bin_file" ]; then
    echo "Error: $bin_file not exist"
    exit
fi

$bin_file "$@"

#/* vim: set ts=4 sts=4 sw=4 et : */
