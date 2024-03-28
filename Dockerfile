FROM golang:1.22

RUN go version
ENV GOPATH=/

COPY ./ ./

# build app
RUN go mod download
RUN go build -o app ./cmd/main.go

CMD ["./app"]