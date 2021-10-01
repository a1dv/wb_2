.PHONY: build
build:
	go vet
	golint
	go build -o orch -v main.go
.DEFAULT_GOAL := build
