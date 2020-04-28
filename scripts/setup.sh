#!/bin/bash
export GOPROXY=https://goproxy.io
export GO111MODULE=on

command=$1
shift

# please make sure running this script using ./scripts/setup.sh

if [[ "$command" == "build" ]]; then
  go build -o build/amazing-proxy src/main.go $@
else
  go run src/main.go $@
fi
