.PHONY: build
build:
	go vet
	golint
	go build -o wget -v main.go
.DEFAULT_GOAL := build

test:
	./wget http://www.gatchina.biz/generator res.txt
	head res.txt
