#!/bin/bash

export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

go build -ldflags="-linkmode=internal -w -s -extldflags '-static' -X main.BuildCPUFlags=native" ../cmd/capitech/main.go

#./updateProd.sh
