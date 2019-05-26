// Code generated by protoc-gen-go. DO NOT EDIT.
// source: setup/code-generation-test/greet/greetpb/greet.proto

package greetpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("setup/code-generation-test/greet/greetpb/greet.proto", fileDescriptor_1c65b8f81d25ac18)
}

var fileDescriptor_1c65b8f81d25ac18 = []byte{
	// 97 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x29, 0x4e, 0x2d, 0x29,
	0x2d, 0xd0, 0x4f, 0xce, 0x4f, 0x49, 0xd5, 0x4d, 0x4f, 0xcd, 0x4b, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc,
	0xcf, 0xd3, 0x2d, 0x49, 0x2d, 0x2e, 0xd1, 0x4f, 0x2f, 0x4a, 0x4d, 0x85, 0x92, 0x05, 0x49, 0x10,
	0x5a, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0x15, 0xcc, 0x31, 0xe2, 0xe3, 0xe2, 0x71, 0x07,
	0x31, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x9d, 0x38, 0xa3, 0xd8, 0xa1, 0xaa, 0x93, 0xd8,
	0xc0, 0x0a, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x45, 0x96, 0xe6, 0x60, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "setup/code-generation-test/greet/greetpb/greet.proto",
}
