# FROM python:3.8-alpine as engine
# COPY src/engine/* .
# RUN pip install flask
# CMD ["python", "engine.py"]

FROM golang:1.21.1 as play-game
COPY src/play-game/* .
WORKDIR /play-game
RUN go mod download
COPY ./*.go ./
RUN go build -o playGame-executable
# CMD ["go", "run", "playGame.go"]
CMD [ "./playGame-executable" ]