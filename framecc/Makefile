all: install

fmt:
	go fmt ./...

grammar.peg.go: pigeon goimports
	pigeon -o grammar.peg.go grammar.peg
	goimports -w grammar.peg.go

get: grammar.peg.go
	go get

pigeon:
	go get github.com/PuerkitoBio/pigeon

goimports:
	go get golang.org/x/tools/cmd/goimports

generate: get
	go generate ./...

install: generate get fmt
	go install

test: install
	go test -v ./...

.PHONY: install generate test get pigeon goimports fmt
