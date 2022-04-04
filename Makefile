VERSION=$(shell cat VERSION)


.PHONY: build_all build_mac build_linux

build_all: build_linux
	zip target/ie-linux-$(VERSION).zip target/ie-linux
	zip target/ie-mac-$(VERSION).zip target/ie-mac

build_mac:
	GOOS=darwin GOARCH=amd64 go build -o target/ie-mac -ldflags="-X 'main.version=$(VERSION)'"

build_linux:
	GOOS=linux GOARCH=amd64 go build -o target/ie-linux -ldflags="-X 'main.version=$(VERSION)'"

