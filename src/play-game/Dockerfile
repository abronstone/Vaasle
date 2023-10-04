FROM golang:1.21.1 as play-game
WORKDIR /play-game

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./*.go ./
RUN go build -o playGame-executable

CMD [ "./playGame-executable" ]