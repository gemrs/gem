all: gem framecc

framecc:
	cd framecc && make

gem: framecc
	cd gem && make

test:
	cd framecc && make test
	cd gem && make test

.PHONY: gem framecc
