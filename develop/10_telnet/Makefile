.PHONY: build
build:
	go vet
	golint
	go build -o telnet -v main.go
.DEFAULT_GOAL := build
