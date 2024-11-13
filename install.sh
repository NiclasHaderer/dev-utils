#!/bin/bash

./build.sh

echo "Installing"
for file in build/*; do
  # Copy the file to the ~/.local/bin directory
  cp "$file" ~/.local/bin
  chmod +x "$HOME/.local/bin/$(basename "$file")"
done