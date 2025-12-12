#!/bin/bash

set -e

PKG_LIST=$(go list ./... | grep -v /vendor/ | xargs) 
go test -covermode=count -coverprofile .coverage $PKG_LIST 

# check if every project files contain the "go-coverage:ignore" comment
for file in $(find . -name '*.go' | grep -v /vendor/); do
    if grep -q '//\s*go-coverage:ignore' $file; then
	    # remove the "./" prefix from the file path
	    file2=${file#"./"}
	    echo "Ignoring coverage for: $file2"

	    # remove the lines of the .coverage file that contains the file2 
	    # value
	    grep -v "${file2}" .coverage > .coverage.tmp && mv .coverage.tmp .coverage
	    
    fi
done

uncover -min 85 .coverage || (rm .coverage && exit 1)
rm .coverage
