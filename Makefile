build:
	go build -o /bin/gpscan cmd/gpscan

fmt:
	go fmt ./...

vet:
	go vet ./...

release:
	goreleaser build