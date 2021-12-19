# syntax=docker/dockerfile:1
FROM golang:1.17-alpine
WORKDIR /GoServer
COPY .. ./
RUN go get -d ./...
RUN go mod tidy
RUN go build -o /docker-go-server
EXPOSE $PORT
CMD ["/docker-go-server","$PORT"]