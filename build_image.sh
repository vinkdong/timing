#! /bin/bash
GOOS=linux GOARCH=amd64 go build timing.go
docker build -t timing .