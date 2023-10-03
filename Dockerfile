FROM python:3.8-alpine as engine
COPY src/engine/* .
RUN pip install flask
CMD ["python", "engine.py"]

FROM python:3.8-alpine as play-game
COPY src/play-game/* .
RUN pip install flask
CMD ["python", "play-game.py"]

# FROM docker as mongo
# RUN docker pull mongo:latest

FROM golang:1.21-bullseye as mongo
COPY src/database/* .
WORKDIR /src/database
RUN go mod init mongo
RUN docker pull mongo:latest
RUN go get go.mongodb.org/mongo-driver/mongo@v1.8.2
CMD ["go", "run", "mongo.go"]
