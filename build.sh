#!/bin/bash

echo "Building"
rm -r build &> /dev/null
mkdir build

go build -o build/repeat cmd/repeat/main.go