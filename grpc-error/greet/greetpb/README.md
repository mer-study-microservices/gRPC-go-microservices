# Greet With Deadline Service

## Service Description
* We'll create a call GreetWithDeadline RPC call
* We'll implement the server to return the response after 3000 ms 
* The server will check if the client has cancelled the request
* We'll implement the client to set a deadline of 5000 ms
* We'll implement the client to set a deadline of 1000 ms

## Generate `.proto` code
```bash
# cd /path/to/GRPC-GO-MICROSERVICES
protoc grpc-error/greet/greetpb/greet.proto --go_out=plugins=grpc:.
```