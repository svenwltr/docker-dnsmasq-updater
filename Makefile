

deps:
	GOPATH=$(shell readlink -f .) go get -d ./...

build:
	GOPATH=$(shell readlink -f .) go build


clean:
	rm -rf src


run:
	GOPATH=$(shell readlink -f .) go run update-dns.go

