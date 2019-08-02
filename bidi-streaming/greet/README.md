# Greet Everyone Service

## GreetEveryone API Definition

* It will take MANY `GreetEveryoneRequest` that contains a `Greeting`
* It will return MANY `GreetEveryoneResponse` that contains a result string

## Generate `.proto` code

```bash
# cd /path/to/GRPC-GO-MICROSERVICES
protoc bidi-streaming/greet/greetpb/greet.proto --go_out=plugins=grpc:.
```