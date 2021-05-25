wire-build:
	./wire_gen.sh

mockgen:
	./mockgen.sh

generate-all: wire-build mockgen

build: generate-all lint test
	docker build . -f build/Dockerfile -t todo:latest

run: build
	docker-compose -f docker-compose/docker-compose.yml up

lint:   ## run go lint on the source files
	golangci-lint run -v

test:
	go test --coverprofile=coverage.out $(go list ./... | grep -v mock)  --race ./...
	go tool cover -func=coverage.out
