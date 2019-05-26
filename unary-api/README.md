# gRPC Unary

Generate `.proto` code

```bash
# cd /path/to/GRPC-GO-MICROSERVICES
protoc unary-api/greet/greetpb/greet.proto --go_out=plugins=grpc:.
```

Server: `greet_server/server.go`

Client: `greet_client/client.go`