// Code generated by protoc-gen-go.
// source: test_grpc.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	test_grpc.proto

It has these top-level messages:
	FooRequest
	FooReply
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FooRequest struct {
}

func (m *FooRequest) Reset()                    { *m = FooRequest{} }
func (m *FooRequest) String() string            { return proto.CompactTextString(m) }
func (*FooRequest) ProtoMessage()               {}
func (*FooRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FooReply struct {
}

func (m *FooReply) Reset()                    { *m = FooReply{} }
func (m *FooReply) String() string            { return proto.CompactTextString(m) }
func (*FooReply) ProtoMessage()               {}
func (*FooReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*FooRequest)(nil), "pb.FooRequest")
	proto.RegisterType((*FooReply)(nil), "pb.FooReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestGrpc service

type TestGrpcClient interface {
	Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooReply, error)
}

type testGrpcClient struct {
	cc *grpc.ClientConn
}

func NewTestGrpcClient(cc *grpc.ClientConn) TestGrpcClient {
	return &testGrpcClient{cc}
}

func (c *testGrpcClient) Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooReply, error) {
	out := new(FooReply)
	err := grpc.Invoke(ctx, "/pb.TestGrpc/Foo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestGrpc service

type TestGrpcServer interface {
	Foo(context.Context, *FooRequest) (*FooReply, error)
}

func RegisterTestGrpcServer(s *grpc.Server, srv TestGrpcServer) {
	s.RegisterService(&_TestGrpc_serviceDesc, srv)
}

func _TestGrpc_Foo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FooRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestGrpcServer).Foo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TestGrpc/Foo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestGrpcServer).Foo(ctx, req.(*FooRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TestGrpc",
	HandlerType: (*TestGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Foo",
			Handler:    _TestGrpc_Foo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test_grpc.proto",
}

func init() { proto.RegisterFile("test_grpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x49, 0x2d, 0x2e,
	0x89, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52,
	0xe2, 0xe1, 0xe2, 0x72, 0xcb, 0xcf, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0xe2, 0xe2,
	0xe2, 0x00, 0xf3, 0x0a, 0x72, 0x2a, 0x8d, 0x0c, 0xb9, 0x38, 0x42, 0x52, 0x8b, 0x4b, 0xdc, 0x8b,
	0x0a, 0x92, 0x85, 0x54, 0xb9, 0x98, 0xdd, 0xf2, 0xf3, 0x85, 0xf8, 0xf4, 0x0a, 0x92, 0xf4, 0x10,
	0xca, 0xa5, 0x78, 0xe0, 0xfc, 0x82, 0x9c, 0x4a, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xb9, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0xba, 0x62, 0xc9, 0x6a, 0x00, 0x00, 0x00,
}