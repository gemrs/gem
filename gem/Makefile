all: install

fmt:
	go fmt ./...

generate: gopygen
	go generate ./...

get:
	go get

gopygen:
	go get github.com/tgascoigne/gopygen

install: generate get fmt
	go install

test: install
	go test -v ./...
	gem `which py.test` -s -v content/

.PHONY: install generate test get gopygen fmt
