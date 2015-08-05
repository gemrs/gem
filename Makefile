all: install

install: generate get fmt
	go install

generate: gopygen
	go generate ./...

test: generate
	go test ./...
	py.test -v content/

get:
	go get

gopygen:
	go get github.com/tgascoigne/gopygen

fmt:
	go fmt ./...

.PHONY: install generate test get gopygen format
