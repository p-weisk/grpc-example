#!/usr/bin/env sh
CGO_ENABLED=0 go build
mv server container/server
