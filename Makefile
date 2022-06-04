build:
	go build -o /bin/gpscan ./cmd

fmt:
	go fmt ./...

vet:
	go vet ./...

update-pkg-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
    go get github.com/$(USER)/$(PACKAGE)@v$(VERSION)

