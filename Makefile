#!/usr/bin/env bash

COVERALLS_TOKEN=L7pI9cMzu1RhGvitE9uwabbCp5ByIvJPR
OUT=c.out

ok:test conver

fmt:
	gofmt -l -w ./

cover:
	go test -cover

file:
	go test -v -covermode=count -coverprofile=${OUT}

func:file
	go tool cover -func=${OUT}

html:file
	go tool cover -html=${OUT} -o ${OUT}.html

goveralls:
	${GOPATH}/bin/goveralls -coverprofile=${OUT} -service=travis-ci -repotoken ${COVERALLS_TOKEN}
