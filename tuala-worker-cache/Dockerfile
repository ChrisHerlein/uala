#BUILD
FROM golang:1.23.0-alpine AS builder

EXPOSE 8080

WORKDIR /go/build

COPY go.mod .
COPY go.sum .
COPY config.go .
COPY engine.go .
COPY handleCache.go .
COPY model.go .
COPY queueHandler.go .
COPY worker.go .

RUN go mod download

RUN GOOS=linux go build -o server

# RUN
FROM alpine:3.15.0
WORKDIR /go/src/app
COPY --from=builder /go/build/server .

CMD ["/go/src/app/server"]
