all: build

build:
	go build -o bin/gendiff cmd/gendiff/main.go