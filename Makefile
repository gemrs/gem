all: gem framecc

framecc:
	cd framecc && make

gem: framecc
	cd gem && make

.PHONY: gem framecc
