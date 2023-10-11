# Engine
FROM golang:1.21.1 as engine

WORKDIR /src/structs
COPY /structs/* ./

WORKDIR /src/engine

COPY ./engine/go.mod ./engine/go.sum ./
RUN go mod download

COPY ./engine/*.go ./
CMD [ "go", "run", "." ]

# Mongo
FROM golang:1.21.1 as mongo

WORKDIR /src/structs
COPY /structs/* ./

WORKDIR /src/mongo
COPY ./mongo/go.mod mongo/go.sum ./

RUN go mod download

COPY ./mongo/*.go ./
CMD [ "go", "run", "." ]

# Play-Game
FROM golang:1.21.1 as play-game

WORKDIR /src/structs
COPY /structs/* ./

WORKDIR /src/play-game

COPY ./play-game/go.mod ./play-game/go.sum ./
RUN go mod download

COPY ./play-game/*.go ./
CMD [ "go", "run", "." ]