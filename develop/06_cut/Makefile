.PHONY: build
build:
	go vet
	golint
	go build -o cut -v main.go
.DEFAULT_GOAL := build
