# syntax=docker/dockerfile:1
FROM golang:1.17-alpine
WORKDIR /GoServer
COPY .. ./
RUN go get -d ./...
RUN go mod tidy
RUN go build -o /docker-go-server
EXPOSE 1200
CMD ["/docker-go-server","1200"]