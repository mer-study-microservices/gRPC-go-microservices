# Square Root Service

## Service Description
* Let's implement an error message for a new SquareRoot Unary API
* We'll create SquareRoot RPC
* We'll implement the Server with the error handling
* We'll implement the Client with the error handling 

## Generate `.proto` code
```bash
# cd /path/to/GRPC-GO-MICROSERVICES
protoc grpc-error/calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
```