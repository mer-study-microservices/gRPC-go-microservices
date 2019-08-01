# Long Greet Service

## LongGreet API Definition

* It will take **MANY** `LongGreetRequest` that contains a `Greeting`
* It will return **ONE** `LongGreetResponse` that contains a result string

## Generate `.proto` code

```bash
# cd /path/to/GRPC-GO-MICROSERVICES
protoc client-streaming/greet/greetpb/greet.proto --go_out=plugins=grpc:.
```