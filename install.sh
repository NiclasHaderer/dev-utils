#!/bin/bash

ARCH="$(uname -m)"

case "${ARCH}" in
	x86_64) ARCH="amd64" ;;
	arm64) ARCH="arm64" ;;
	aarch64) ARCH="arm64" ;;
	*)
		echo "Unsupported architecture: ${ARCH}"
		exit 2
		;;
esac

OS="$(uname -s)"
case "${OS}" in
	Linux) OS="linux" ;;
	Darwin) OS="darwin" ;;
	*)
		echo "Unsupported OS: ${OS}"
		exit 3
		;;
esac


./compile.sh "$1" "$OS" "$ARCH"
if [ $? -ne 0 ]; then
  exit 1
fi

echo "Installing"
for file in bin/*"$OS-$ARCH"; do
  # Copy the file to the ~/.local/bin directory
  command_name=$(basename "$file")
  command_name=${command_name%"-$OS-$ARCH"}
  destination=~/.local/bin/"$command_name"
  cp "$file" "$destination"
  chmod +x "$destination"
done