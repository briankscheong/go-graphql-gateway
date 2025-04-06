FROM golang:1.23-alpine as go-builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build server.go

FROM alpine:3.21

COPY --from=go-builder /app/server ./server

EXPOSE 8080
EXPOSE 50051

CMD ["./server"]
