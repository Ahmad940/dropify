#!/bin/bash

# build my source code
go build -o ./build

# Check if the directory exists
if [[ ! -d "/root/dropify" ]]; then
  # Create the directory if it does not exists
  mkdir -p "/root/dropify"
fi

# move the built file
mv ../dist/build /root/dropify