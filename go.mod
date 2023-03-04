module github.com/akuity/protoc-gen-grpc-gateway-client

go 1.20

require (
	github.com/akuity/go-akuity v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	github.com/stretchr/testify v1.8.2
	google.golang.org/genproto v0.0.0-20230303212802-e74f57abe488
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.2-0.20230222093303-bc1253ad3743
)

require (
	github.com/alevinval/sse v1.0.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/akuity/go-akuity => ../go-akuity
