#!/bin/bash

# Get version from git
VERSION=$1

# Check if version is set
if [ -z "$VERSION" ]; then
	echo "Version not set"
	exit 1
fi

# Variables
OUTPUT_DIR="bin"
GOOS_ARRAY=("darwin" "darwin" "linux" "linux" "windows" "windows")
GOARCH_ARRAY=("amd64" "arm64" "amd64" "arm64" "amd64" "arm64")

ONLY_OS=$2
ONLY_ARCH=$3

if [ -n "$ONLY_OS" ]; then
  GOOS_ARRAY=("$ONLY_OS")
fi

if [ -n "$ONLY_ARCH" ]; then
  GOARCH_ARRAY=("$ONLY_ARCH")
fi


# Create output directory
rm -r $OUTPUT_DIR
mkdir -p $OUTPUT_DIR

# Loop through architectures
for cmd in cmd/*; do
  BINARY_NAME=$(basename "$cmd")
  # shellcheck disable=SC2068
  # Iterate over all folders in cmd
  for i in ${!GOOS_ARRAY[@]}; do
    GOOS=${GOOS_ARRAY[$i]}
    GOARCH=${GOARCH_ARRAY[$i]}

    OUTPUT_NAME="$OUTPUT_DIR/$BINARY_NAME-$GOOS-$GOARCH"

    # Cross-compile
    env GOOS="$GOOS" GOARCH="$GOARCH" go build -ldflags "-X 'dev-utils/lib/config.Version=$VERSION'" -o "$OUTPUT_NAME" "$cmd/main.go"

    # Check if cross-compilation was successful
    if [ $? -eq 0 ]; then
      echo "Successfully compiled $BINARY_NAME for $GOOS $GOARCH"
    else
      echo "Error compiling $BINARY_NAME for $GOOS $GOARCH"
      exit 1
    fi
  done

done
