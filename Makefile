.PHONY: build fmt test vet run

GOPATH := ${PWD}

default: build
build:
	GOPATH=$(GOPATH) go build -v -o ./bin/goco ./
run:
	GOPATH=$(GOPATH) go run ./goco.go -type=test -config=./testdata/test.json -out=./testdata/test.go
vet:
	GOPATH=$(GOPATH) go vet ./...
fmt:
	GOPATH=$(GOPATH) go fmt ./...
test:
	GOPATH=$(GOPATH) go test ./...