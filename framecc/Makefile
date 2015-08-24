all: install

fmt:
	go fmt ./...

get: 
	go get

generate: get
	go generate ./...

install: generate get fmt
	go install

test: install
	go test -v ./...

.PHONY: install generate test get pigeon goimports fmt
