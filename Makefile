all: build

build:
	go build -o bin/gendiff cmd/gendiff/main.go

test:
	go test -v

lint:
	golangci-lint run

coverage:
	go test -coverprofile=coverage.out ./...
