#BUILD
FROM golang:1.23.0-alpine AS builder

EXPOSE 8080

WORKDIR /go/build

COPY go.mod .
COPY go.sum .
COPY api.go .
RUN go mod download

COPY ./src src

RUN GOOS=linux go build -o server ./api.go

# RUN
FROM alpine:3.15.0
WORKDIR /go/src/app
COPY --from=builder /go/build/server .

CMD ["/go/src/app/server"]
