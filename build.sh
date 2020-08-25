#!/bin/bash

# This script builds a binary for whatever OS you're running it on

go test
TEST_RESULT=`echo $?`
if test $TEST_RESULT -ne 0; then
  echo "tests failed. Exiting"
  exit $TEST_RESULT
fi

VERSION=`cat version`

mkdir -p bin/$VERSION

go build -o bin/$VERSION/logspam ./cmd
