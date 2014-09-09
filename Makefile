

build:
	make deps
	GOPATH=$(shell readlink -f .) go build

deps:
	GOPATH=$(shell readlink -f .) go get -d ./...

clean:
	rm -rf src


run:
	GOPATH=$(shell readlink -f .) go run update-dnsmasq-updater.go

