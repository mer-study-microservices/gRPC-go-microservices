# Client Setup Boilerplate Code

1. Generate `.proto` code

```bash
# cd /path/to/GRPC-GO-MICROSERVICES
protoc setup/client-boilerplate/greet/greetpb/greet.proto --go_out=plugins=grpc:.
```

2. The Client Setup Boilerplate Code is in `/setup/server-boilerplate/greet_client/client.go`