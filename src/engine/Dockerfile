FROM golang:1.21.1 as engine 
WORKDIR /engine

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./*.go ./
RUN go build -o engine-executable

CMD [ "./engine-executable" ]