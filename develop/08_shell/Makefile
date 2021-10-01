.PHONY: build
build:
	go vet
	golint
	go build -o shell -v main.go
.DEFAULT_GOAL := build
