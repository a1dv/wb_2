.PHONY: build
build:
	go vet
	golint
	go build -o anagrams -v main.go
.DEFAULT_GOAL := build
