FROM golang:1.15.2-alpine3.12

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main ./cmd/server
EXPOSE 8081
CMD ["/app/main"]