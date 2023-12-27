package main

import (
	"context"
	"fmt"
	"os"

	"github.com/omrikiei/grpc-gateway-client/pkg/grpc/gateway"

	grpc_gateway_client_example "github.com/akuity/grpc-gateway-client-example/pkg/api/gen"
)

func main() {
	baseURL := "http://localhost:8081"
	if len(os.Args) > 1 {
		baseURL = os.Args[1]
	}
	client := grpc_gateway_client_example.NewGreeterGatewayClient(gateway.NewClient(baseURL))
	resp, err := client.SayHello(context.Background(), &grpc_gateway_client_example.HelloRequest{Name: "World"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Message)
}
