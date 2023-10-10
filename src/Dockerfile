# Engine
FROM golang:1.21.1 as engine

WORKDIR /src/structs
COPY /structs/* ./

WORKDIR /src/engine

COPY ./engine/go.mod ./engine/go.sum ./
RUN go mod download

COPY ./engine/*.go ./
RUN go build -o engine-executable

CMD [ "./engine-executable" ]

# Mongo
FROM golang:1.21.1 as mongo

WORKDIR /mongo
COPY ./mongo/go.mod mongo/go.sum ./

RUN go mod download
COPY ./mongo/*.go ./

RUN go build -o mongo-executable
CMD [ "./mongo-executable" ]

# Play-Game
FROM golang:1.21.1 as play-game
WORKDIR /play-game

COPY ./play-game/go.mod ./play-game/go.sum ./
RUN go mod download

COPY ./play-game/*.go ./
RUN go build -o playGame-executable

CMD [ "./playGame-executable" ]