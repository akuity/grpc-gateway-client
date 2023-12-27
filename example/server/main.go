package main

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gen "github.com/omrikiei/grpc-gateway-client-example/pkg/api/gen"
)

type greeter struct {
	gen.UnimplementedGreeterServer
}

func (g *greeter) SayHello(_ context.Context, req *gen.HelloRequest) (*gen.HelloReply, error) {
	return &gen.HelloReply{Message: "Hello " + req.Name}, nil
}

func main() {
	ctx := context.TODO()
	srv := grpc.NewServer()
	gen.RegisterGreeterServer(srv, &greeter{})

	grpcListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := gen.RegisterGreeterHandlerFromEndpoint(ctx, mux, "localhost:8080", opts); err != nil {
		panic(err)
	}

	go srv.Serve(grpcListener)
	go http.ListenAndServe(":8081", mux)
	println("Listening on :8081")
	<-ctx.Done()
}
