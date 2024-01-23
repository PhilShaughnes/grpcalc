FROM golang:1.21-alpine3.19 AS builder

WORKDIR /app

ADD . /app

RUN go build -o bin/server server/main.go

FROM alpine:3.19
# FROM scratch

COPY --from=builder /app/bin/server .

CMD ["./server"]

