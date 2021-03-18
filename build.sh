#!/bin/bash

# This script builds a binary for whatever OS you're running from

if ! go test; then
  echo "tests failed. Exiting"
  exit 1
fi

go build -o bin/logspam .
