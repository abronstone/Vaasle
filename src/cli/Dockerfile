# Adapted from smaple Dockerfile on Go image page https://hub.docker.com/_/golang

FROM golang:1.21 AS cli

WORKDIR /usr/src/cli

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/cli ./...

CMD ["cli"]