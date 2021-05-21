wire-build:
	wire ./internal/server

mockgen:
	./mockgen.sh

generate-all: wire-build mockgen

build-local-server: generate-all
	go build -ldflags="-X main.ApplicationVersion=local" -o bin/app ./cmd/cp-provision

run-server-with-nats: wire-build
	docker-compose up -d
	go run ./cmd/cp-provision serve

run-server: wire-build
	go run ./cmd/cp-provision serve

lint:   ## run go lint on the source files
	golangci-lint run -v
