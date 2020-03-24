#!/bin/bash
set -e

echo "#################### downloading CompileDaemon"
# disable go modules to avoid this package from getting into go.mod
# as we only need it locally to watch and rebuild server on change
GO111MODULE=off go get github.com/githubnemo/CompileDaemon

echo "#################### starting deamon"
CompileDaemon --build="go build -o main cmd/api/main.go" --command=./main
