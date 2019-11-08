#!/bin/bash

LICENSE=$(cat <<EOF
// Copyright 2019 The fsn-go-sdk Authors
// This file is part of the fsn-go-sdk library.
//
// The fsn-go-sdk library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The fsn-go-sdk library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the fsn-go-sdk library. If not, see <http://www.gnu.org/licenses/>.

EOF
)

if (( $# > 0 )) && [[ "$1" = "-h" || "$1" = "--help" ]]; then
    echo "Usage: $0 <file>"
    exit
fi

for file in $@; do
    [[ ! -f "$file" ]] && echo "ignore non-exist file $file" && continue

    # already has the license
    if grep -q "GNU Lesser General Public License" $file; then
        echo "ignore $file because it already has the license"
        continue
    fi

    echo -e "$LICENSE\n" | cat - $file > /tmp/tempfile
    mv /tmp/tempfile $file
    echo "add license to $file succeed"
done

#/* vim: set ts=4 sts=4 sw=4 et : */
