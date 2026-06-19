.DEFAULT_GOAL := build
.PHONY:fmt vet fix build
fmt:
	go fmt ./...
vet: fmt
	go vet ./...
fix: vet
	go fix ./...
build: fix
	go build -o main.out
