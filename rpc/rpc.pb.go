// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	HelloRequest
	Picture
	HelloReply
	Request
*/
package rpc

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

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Picture struct {
	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Type  []byte `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *Picture) Reset()                    { *m = Picture{} }
func (m *Picture) String() string            { return proto.CompactTextString(m) }
func (*Picture) ProtoMessage()               {}
func (*Picture) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Picture) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Picture) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Request struct {
	Filename string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Request) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "rpc.HelloRequest")
	proto.RegisterType((*Picture)(nil), "rpc.Picture")
	proto.RegisterType((*HelloReply)(nil), "rpc.HelloReply")
	proto.RegisterType((*Request)(nil), "rpc.Request")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FileServer service

type FileServerClient interface {
	// Sends a greeting
	GetImage(ctx context.Context, in *Request, opts ...grpc.CallOption) (FileServer_GetImageClient, error)
}

type fileServerClient struct {
	cc *grpc.ClientConn
}

func NewFileServerClient(cc *grpc.ClientConn) FileServerClient {
	return &fileServerClient{cc}
}

func (c *fileServerClient) GetImage(ctx context.Context, in *Request, opts ...grpc.CallOption) (FileServer_GetImageClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FileServer_serviceDesc.Streams[0], c.cc, "/rpc.FileServer/GetImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServerGetImageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileServer_GetImageClient interface {
	Recv() (*Picture, error)
	grpc.ClientStream
}

type fileServerGetImageClient struct {
	grpc.ClientStream
}

func (x *fileServerGetImageClient) Recv() (*Picture, error) {
	m := new(Picture)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for FileServer service

type FileServerServer interface {
	// Sends a greeting
	GetImage(*Request, FileServer_GetImageServer) error
}

func RegisterFileServerServer(s *grpc.Server, srv FileServerServer) {
	s.RegisterService(&_FileServer_serviceDesc, srv)
}

func _FileServer_GetImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServerServer).GetImage(m, &fileServerGetImageServer{stream})
}

type FileServer_GetImageServer interface {
	Send(*Picture) error
	grpc.ServerStream
}

type fileServerGetImageServer struct {
	grpc.ServerStream
}

func (x *fileServerGetImageServer) Send(m *Picture) error {
	return x.ServerStream.SendMsg(m)
}

var _FileServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.FileServer",
	HandlerType: (*FileServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetImage",
			Handler:       _FileServer_GetImage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x4b, 0xc5, 0x30,
	0x14, 0xc5, 0x8d, 0x5f, 0x7d, 0xbd, 0x14, 0x84, 0xe0, 0x50, 0x1e, 0x0e, 0x12, 0x50, 0xc4, 0x21,
	0xa8, 0x5d, 0x9c, 0x3b, 0xf8, 0xb1, 0x95, 0x38, 0x38, 0xd7, 0x7a, 0x6d, 0x03, 0x89, 0x89, 0x49,
	0xaa, 0xf6, 0xbf, 0x97, 0xa4, 0xad, 0xbc, 0xed, 0xfc, 0x6e, 0xb8, 0xf7, 0x9c, 0x1c, 0xc8, 0x9d,
	0xed, 0xb8, 0x75, 0x26, 0x18, 0x7a, 0xe0, 0x6c, 0xc7, 0x18, 0x14, 0x4f, 0xa8, 0x94, 0x11, 0xf8,
	0x35, 0xa2, 0x0f, 0x94, 0xc2, 0xe1, 0x67, 0xab, 0xb1, 0x24, 0xe7, 0xe4, 0x2a, 0x17, 0x49, 0xb3,
	0x0a, 0xb2, 0x46, 0x76, 0x61, 0x74, 0x48, 0x4f, 0xe1, 0x48, 0xea, 0xb6, 0x9f, 0xdf, 0x0b, 0x31,
	0x43, 0x5c, 0x0a, 0x93, 0xc5, 0x72, 0x3f, 0x0d, 0x93, 0x66, 0x97, 0x00, 0xcb, 0x61, 0xab, 0x26,
	0x5a, 0x42, 0xa6, 0xd1, 0xfb, 0x75, 0x33, 0x17, 0x2b, 0xb2, 0x0b, 0xc8, 0x56, 0xef, 0x2d, 0x6c,
	0x3e, 0xa4, 0xc2, 0x1d, 0xff, 0x7f, 0xbe, 0xbb, 0x07, 0x78, 0x90, 0x0a, 0x5f, 0xd0, 0x7d, 0xa3,
	0xa3, 0xd7, 0xb0, 0x79, 0xc4, 0xf0, 0x9c, 0xcc, 0x0b, 0x1e, 0xbf, 0xb4, 0xdc, 0xd8, 0xce, 0xb4,
	0xc4, 0x65, 0x7b, 0x37, 0xa4, 0xbe, 0x85, 0x33, 0x69, 0x78, 0x3f, 0xc4, 0x39, 0xfe, 0xb6, 0xda,
	0x2a, 0xf4, 0x7c, 0x88, 0xd1, 0x7e, 0x8c, 0x53, 0xef, 0xf5, 0x49, 0x8a, 0xf9, 0x1a, 0x75, 0x13,
	0x7b, 0x69, 0xc8, 0xdb, 0x71, 0x2a, 0xa8, 0xfa, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x12, 0xce, 0x40,
	0x9d, 0x2d, 0x01, 0x00, 0x00,
}