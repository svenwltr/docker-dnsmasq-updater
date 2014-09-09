

build:
	make deps
	GOPATH=$(shell readlink -f .) go build -o docker-dnsmasq-updater

deps:
	GOPATH=$(shell readlink -f .) go get -d ./...

clean:
	rm -rf src


run:
	GOPATH=$(shell readlink -f .) go run docker-dnsmasq-updater.go

