.PHONY: generate-test

generate-test:
	mkdir -p ./bin
	go build -o bin/protoc-gen-grpc-gateway-client
	buf generate
