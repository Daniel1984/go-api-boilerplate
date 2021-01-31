#!/bin/bash
set -e

echo "#################### migrating dev db"
go run cmd/dbmigrate/main.go

echo "#################### migrating test db"
go run cmd/dbmigrate/main.go -dbname=boilerplatetest

echo "################### removing the .git dir"

rm -rf .git

echo "#################### downloading CompileDaemon"
# disable go modules to avoid this package from getting into go.mod
# as we only need it locally to watch and rebuild server on change
GO111MODULE=off go get github.com/githubnemo/CompileDaemon

echo "#################### starting deamon"
CompileDaemon --build="go build -o main cmd/api/main.go" --command=./main
