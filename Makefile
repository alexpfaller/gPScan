build:
	go build -o /bin/gpscan .

fmt:
	go fmt ./...

vet:
	go vet ./...
