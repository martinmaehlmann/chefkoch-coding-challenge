FROM golang:alpine as builder
WORKDIR /app/
COPY . .
RUN go build -o bin/app ./cmd/todo
FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/bin/app /app
ENTRYPOINT ["./app"]