FROM golang:1.21.1 as engine 
WORKDIR /engine

COPY src/engine/go.mod src/engine/go.sum ./
RUN go mod download

COPY src/engine/*.go ./
RUN go build -o engine-executable

CMD [ "./engine-executable" ]