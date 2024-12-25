CC=go
VERSION=$(shell cat VERSION)
NAME_BASE=noerrorcode
SOURCES=main.go \
		client.go \
		service.go \
		local_client.go \
		config.go \
		websocket.go \
		vars.go

APP=$(NAME_BASE)

all: linux darwin windows
linux: directories
linux: bin/linux/$(APP)
darwin: directories
darwin: bin/darwin/$(APP)
windows: directories
windows: bin/windows/$(APP).exe

bin/linux/$(APP): $(SOURCES)
	GOOS=linux GOARCH=amd64 $(CC) build -ldflags="-w -s -X main.AppVersion=$(VERSION)" -o $@ -v $^

bin/darwin/$(APP): $(SOURCES)
	GOOS=darwin $(CC) build -ldflags="-w -s -X main.AppVersion=$(VERSION)" -o $@ -v $^

bin/windows/$(APP).exe: $(SOURCES)
	GOOS=windows $(CC) build -ldflags="-w -s -X main.AppVersion=$(VERSION)" -o $@ -v $^

clean:
	-rm -f bin/linux/$(APP)*
	-rm -f bin/darwin/$(APP)*
	-rm -f bin/windows/$(APP)*

distclean:
	-rm -rf bin

tests:
	go clean -testcache
	go test -v ./...

cover:
	@go test -cover | grep coverage:

directories:
	@mkdir -p bin
	@mkdir -p bin/linux
	@mkdir -p bin/darwin
	@mkdir -p bin/windows

run:
	go run server.go
