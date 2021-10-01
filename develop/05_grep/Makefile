.PHONY: build
build:
	go vet
	golint
	go build -o grep -v main.go
.DEFAULT_GOAL := build
