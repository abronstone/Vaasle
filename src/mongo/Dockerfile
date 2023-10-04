# FROM golang:1.21.1 as mongo
# WORKDIR /mongo
# COPY ./go.mod ./go.sum ./
# # WORKDIR /src/database
# RUN go mod download
# COPY ./*.go ./

# # RUN go get go.mongodb.org/mongo-driver/mongo@v1.8.2
# RUN go build -o mongo-executable
# CMD [ "./mongo-executable" ]

FROM mongo:latest as mongo
COPY ./data ./