// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	types.proto

It has these top-level messages:
	NewIDRequest
	NewIDResponse
*/
package types

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

type NewIDRequest struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
}

func (m *NewIDRequest) Reset()                    { *m = NewIDRequest{} }
func (m *NewIDRequest) String() string            { return proto.CompactTextString(m) }
func (*NewIDRequest) ProtoMessage()               {}
func (*NewIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NewIDRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type NewIDResponse struct {
	Name string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	ID   uint64 `protobuf:"varint,3,opt,name=ID" json:"ID,omitempty"`
}

func (m *NewIDResponse) Reset()                    { *m = NewIDResponse{} }
func (m *NewIDResponse) String() string            { return proto.CompactTextString(m) }
func (*NewIDResponse) ProtoMessage()               {}
func (*NewIDResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NewIDResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NewIDResponse) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto.RegisterType((*NewIDRequest)(nil), "types.NewIDRequest")
	proto.RegisterType((*NewIDResponse)(nil), "types.NewIDResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Chief service

type ChiefClient interface {
	NewID(ctx context.Context, in *NewIDRequest, opts ...grpc.CallOption) (*NewIDResponse, error)
}

type chiefClient struct {
	cc *grpc.ClientConn
}

func NewChiefClient(cc *grpc.ClientConn) ChiefClient {
	return &chiefClient{cc}
}

func (c *chiefClient) NewID(ctx context.Context, in *NewIDRequest, opts ...grpc.CallOption) (*NewIDResponse, error) {
	out := new(NewIDResponse)
	err := grpc.Invoke(ctx, "/types.Chief/NewID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Chief service

type ChiefServer interface {
	NewID(context.Context, *NewIDRequest) (*NewIDResponse, error)
}

func RegisterChiefServer(s *grpc.Server, srv ChiefServer) {
	s.RegisterService(&_Chief_serviceDesc, srv)
}

func _Chief_NewID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChiefServer).NewID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Chief/NewID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChiefServer).NewID(ctx, req.(*NewIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chief_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.Chief",
	HandlerType: (*ChiefServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewID",
			Handler:    _Chief_NewID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "types.proto",
}

func init() { proto.RegisterFile("types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x94, 0xb8, 0x78, 0xfc,
	0x52, 0xcb, 0x3d, 0x5d, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xfc,
	0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x63, 0x2e, 0x5e,
	0xa8, 0x9a, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xb8, 0x22, 0x26, 0x84, 0x22, 0x21, 0x3e, 0x2e,
	0x26, 0x4f, 0x17, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x26, 0x4f, 0x17, 0x23, 0x6b, 0x2e,
	0x56, 0xe7, 0x8c, 0xcc, 0xd4, 0x34, 0x21, 0x23, 0x2e, 0x56, 0xb0, 0x6e, 0x21, 0x61, 0x3d, 0x88,
	0xfd, 0xc8, 0xf6, 0x49, 0x89, 0xa0, 0x0a, 0x42, 0x2c, 0x48, 0x62, 0x03, 0xbb, 0xd1, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x17, 0x9e, 0x33, 0x88, 0xb2, 0x00, 0x00, 0x00,
}
