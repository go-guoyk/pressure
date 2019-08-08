#!/bin/bash

set -eu

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build -o pressure

docker build -t guoyk/pressure .
