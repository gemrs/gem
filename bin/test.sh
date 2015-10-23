#!/bin/bash
go test -tags test_python -coverpkg=./gem/... -covermode=count -coverprofile=python.coverprofile ./gem/cmd/gem
go test -tags test_python -covermode=count -coverprofile=gem_log.coverprofile ./gem/log
overalls -project=github.com/sinusoids/gem/gem -covermode=count -debug -ignore=".git,vendor"
overalls -project=github.com/sinusoids/gem/bbc -covermode=count -debug -ignore=".git,vendor"
gocovmerge gem/overalls.coverprofile bbc/overalls.coverprofile python.coverprofile gem_log.coverprofile > coverage.profile
