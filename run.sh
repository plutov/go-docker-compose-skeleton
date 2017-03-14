#!/bin/bash

cd /go/src/github.com/plutov/go-docker-compose-skeleton

glide install
goose --env=development up

go build -o /go/bin/api cmd/api/main.go
go generate cmd/api/main.go

/go/bin/api --addr=":8080" --db="app:12345@(db:3306)/app" --queue="queue:11300"