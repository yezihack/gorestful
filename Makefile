#!/usr/bin/env bash

COVERALLS_TOKEN=L7pI9cMzu1RhGvitE9uwabbCp5ByIvJPR

ok:test conver

test:
	go test -v -covermode=count -coverprofile=coverage.out

conver:
	${GOPATH}/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken ${COVERALLS_TOKEN}
