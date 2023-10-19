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

WORKDIR /src/structs
COPY /structs/* ./

WORKDIR /src/mongo
COPY ./mongo/go.mod mongo/go.sum ./

RUN go mod download
COPY ./mongo/*.go ./

RUN go build -o mongo-executable
CMD [ "./mongo-executable" ]


# Gateway
FROM golang:1.21.1 as gateway

WORKDIR /src/structs
COPY /structs/* ./

WORKDIR /src/gateway

COPY ./gateway/go.mod ./gateway/go.sum ./
RUN go mod download

COPY ./gateway/*.go ./
RUN go build -o gateway-executable

CMD [ "./gateway-executable" ]

# CLI
FROM golang:1.21.1 as cli

WORKDIR /src/cli

COPY ./cli/go.mod ./
RUN go mod download

COPY ./cli/*.go ./
RUN go build -o cli-executable

CMD [ "./cli-executable" ]

# Front End
FROM node:16 AS frontend

WORKDIR /src/frontend

COPY ./frontend/package.json ./

RUN npm install

COPY . ./

EXPOSE 3000

CMD ["npm", "run", "start"]