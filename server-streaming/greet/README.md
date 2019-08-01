# Greet Many Times Service

## GreetManyTimes API Definition 

* It will take **`ONE`** `GreetManyTimesRequest` that contains a `Greeting`
* It will return **`MANY`** `GreetManyTimesResponse` that contains a result string

## Generate `.proto` code

```bash
# cd /path/to/GRPC-GO-MICROSERVICES
protoc server-streaming/greet/greetpb/greet.proto --go_out=plugins=grpc:.
```