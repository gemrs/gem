all: install

install: generate get fmt
	go install

generate: gopygen
	go generate ./...

test: install
	go test ./...
	gem `which py.test` -s -v content/

get:
	go get

gopygen:
	go get github.com/tgascoigne/gopygen

fmt:
	go fmt ./...

.PHONY: install generate test get gopygen format
