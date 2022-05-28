FROM golang:1.17-alpine

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -v -o /bin/gpscan cmd/gpscan/main.go

ENTRYPOINT ["/bin/gpscan"]

CMD [ "/bin/gpscan" ]