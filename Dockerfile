FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . /app/
RUN go mod download
RUN go build -o server cmd/grpc_server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 50051

CMD ["/app/server"]