.PHONY: build
build:
	go vet
	golint
	go build -o server -v main.go
.DEFAULT_GOAL := build
