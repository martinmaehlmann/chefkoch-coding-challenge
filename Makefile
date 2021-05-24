wire-build:
	./wire_gen.sh

mockgen:
	./mockgen.sh

generate-all: wire-build mockgen

build: generate-all
	docker build . -f build/Dockerfile -t todo:latest

run-server-with-nats: wire-build
	go run ./cmd/cp-provision serve

run-server: wire-build
	go run ./cmd/cp-provision serve

lint:   ## run go lint on the source files
	golangci-lint run -v
