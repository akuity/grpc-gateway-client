# grpc-gateway-client

The `grpc-gateway-client` is a high quality REST client generator for [gRPC](https://grpc.io/) services that are fronted by [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway).

## Features

- Strongly typed client interface.
- Supports all gRPC features including streaming.
- Supports all grpc-gateway features including custom query parameters, and request body.
- Battle tested by [Akuity's](https://akuity.io/) production services.


## Usage

1. Install `grpc-gateway-client`:

    ```bash
    $ go install github.com/akuity/grpc-gateway-client/protoc-gen-grpc-gateway-client@latest
    ```
1. Add plugin in your buf.gen.yaml:

    ```yaml
    version: v1
    managed:
    enabled: true

    plugins:
      - name: gateway-client
        path: protoc-gen-grpc-gateway-client
        out: pkg/api/gen
        opt:
        - paths=source_relative
    ```
1. Generate client using `buf generate` and use it in your code:

   ```go
    client := grpc_gateway_client_example.NewGreeterGatewayClient(gateway.NewClient(baseURL))
    resp, err := client.SayHello(context.Background(), &grpc_gateway_client_example.HelloRequest{Name: "World"})
    if err != nil {
        panic(err)
    }
    println(resp.Message)
    ```

See [example](./example/README.md) for a complete example.