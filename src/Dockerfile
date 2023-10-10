FROM golang:1.21.1 as engine

WORKDIR /src/structs
COPY /structs/* ./

WORKDIR /src/engine

COPY ./engine/go.mod ./engine/go.sum ./
RUN go mod download

COPY ./engine/*.go ./
RUN go build -o engine-executable

CMD [ "./engine-executable" ]