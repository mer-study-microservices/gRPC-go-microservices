# Sum API

Generate `.proto` code

```bash
# cd /path/to/GRPC-GO-MICROSERVICES
protoc unary-api/calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
```