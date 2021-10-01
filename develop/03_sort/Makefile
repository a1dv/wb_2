.PHONY: build
build:
	go vet
	golint
	go build -o sort -v main.go
.DEFAULT_GOAL := build
