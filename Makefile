.PHONY: build generate test

build:
	mkdir -p ./bin
	go build -o bin/protoc-gen-grpc-gateway-client ./protoc-gen-grpc-gateway-client

generate:
	buf generate

test: build generate
	go test ./...
