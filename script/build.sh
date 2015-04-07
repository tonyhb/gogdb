#!/bin/bash

go build -o bin main.go
"$GOPATH/bin/golang-nw-pkg" -app=./bin -name="gogdb" -bin="gogdb" -toolbar=true
rm ./bin
