.PHONY: all clean

include .env
export

YARN=$(shell which yarn)
BINARY=invoicer
BUILD=$(shell git rev-parse --short HEAD)
LDFLAGS=-ldflags "-X main.Version=$(APP_VERSION) -X main.Build=$(BUILD)"

all:
	$(YARN) --cwd ./ui install && $(YARN) --cwd ./ui build
	go build -o $(BINARY) $(LDFLAGS) cmd/*

clean:
	-rm -rf ./ui/node_modules
	-rm $(BINARY)
