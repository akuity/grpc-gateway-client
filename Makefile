.PHONY: test

test:
	mkdir -p ./bin
	go build -o bin/protoc-gen-grpc-gateway-client
	buf generate
	go test ./...
