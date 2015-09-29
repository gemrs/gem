#!/bin/bash
cd src
go get github.com/wadey/gocovmerge
packages=`find gem bbc -type f -name "*.go" -printf "%h\n" | sort | uniq | grep -v vendor`
covfiles=""
for f in $packages; do
  go test -v -coverprofile $f/coverage.profile $f
  [ -f $f/coverage.profile ] && covfiles+="$f/coverage.profile "
done
gocovmerge $covfiles > ../coverage.profile
cd -
