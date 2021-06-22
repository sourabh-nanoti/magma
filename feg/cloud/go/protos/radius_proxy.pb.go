// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feg/protos/radius_proxy.proto

package protos // import "magma/feg/cloud/go/protos"

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

type AaaRequest struct {
	Packet               []byte   `protobuf:"bytes,1,opt,name=packet,proto3" json:"packet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AaaRequest) Reset()         { *m = AaaRequest{} }
func (m *AaaRequest) String() string { return proto.CompactTextString(m) }
func (*AaaRequest) ProtoMessage()    {}
func (*AaaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_radius_proxy_66edf419a1df974c, []int{0}
}
func (m *AaaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AaaRequest.Unmarshal(m, b)
}
func (m *AaaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AaaRequest.Marshal(b, m, deterministic)
}
func (dst *AaaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AaaRequest.Merge(dst, src)
}
func (m *AaaRequest) XXX_Size() int {
	return xxx_messageInfo_AaaRequest.Size(m)
}
func (m *AaaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AaaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AaaRequest proto.InternalMessageInfo

func (m *AaaRequest) GetPacket() []byte {
	if m != nil {
		return m.Packet
	}
	return nil
}

type AaaResponse struct {
	Packet               []byte   `protobuf:"bytes,1,opt,name=packet,proto3" json:"packet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AaaResponse) Reset()         { *m = AaaResponse{} }
func (m *AaaResponse) String() string { return proto.CompactTextString(m) }
func (*AaaResponse) ProtoMessage()    {}
func (*AaaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_radius_proxy_66edf419a1df974c, []int{1}
}
func (m *AaaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AaaResponse.Unmarshal(m, b)
}
func (m *AaaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AaaResponse.Marshal(b, m, deterministic)
}
func (dst *AaaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AaaResponse.Merge(dst, src)
}
func (m *AaaResponse) XXX_Size() int {
	return xxx_messageInfo_AaaResponse.Size(m)
}
func (m *AaaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AaaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AaaResponse proto.InternalMessageInfo

func (m *AaaResponse) GetPacket() []byte {
	if m != nil {
		return m.Packet
	}
	return nil
}

func init() {
	proto.RegisterType((*AaaRequest)(nil), "magma.feg.AaaRequest")
	proto.RegisterType((*AaaResponse)(nil), "magma.feg.AaaResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RadiusProxyClient is the client API for RadiusProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RadiusProxyClient interface {
	ProxyPacket(ctx context.Context, in *AaaRequest, opts ...grpc.CallOption) (*AaaResponse, error)
}

type radiusProxyClient struct {
	cc *grpc.ClientConn
}

func NewRadiusProxyClient(cc *grpc.ClientConn) RadiusProxyClient {
	return &radiusProxyClient{cc}
}

func (c *radiusProxyClient) ProxyPacket(ctx context.Context, in *AaaRequest, opts ...grpc.CallOption) (*AaaResponse, error) {
	out := new(AaaResponse)
	err := c.cc.Invoke(ctx, "/magma.feg.RadiusProxy/ProxyPacket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RadiusProxyServer is the server API for RadiusProxy service.
type RadiusProxyServer interface {
	ProxyPacket(context.Context, *AaaRequest) (*AaaResponse, error)
}

func RegisterRadiusProxyServer(s *grpc.Server, srv RadiusProxyServer) {
	s.RegisterService(&_RadiusProxy_serviceDesc, srv)
}

func _RadiusProxy_ProxyPacket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AaaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RadiusProxyServer).ProxyPacket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.RadiusProxy/ProxyPacket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RadiusProxyServer).ProxyPacket(ctx, req.(*AaaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RadiusProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.RadiusProxy",
	HandlerType: (*RadiusProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProxyPacket",
			Handler:    _RadiusProxy_ProxyPacket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feg/protos/radius_proxy.proto",
}

func init() {
	proto.RegisterFile("feg/protos/radius_proxy.proto", fileDescriptor_radius_proxy_66edf419a1df974c)
}

var fileDescriptor_radius_proxy_66edf419a1df974c = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x4b, 0x4d, 0xd7,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f, 0xd6, 0x2f, 0x4a, 0x4c, 0xc9, 0x2c, 0x2d, 0x8e, 0x2f, 0x28,
	0xca, 0xaf, 0xa8, 0xd4, 0x03, 0x8b, 0x09, 0x71, 0xe6, 0x26, 0xa6, 0xe7, 0x26, 0xea, 0xa5, 0xa5,
	0xa6, 0x2b, 0xa9, 0x70, 0x71, 0x39, 0x26, 0x26, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08,
	0x89, 0x71, 0xb1, 0x15, 0x24, 0x26, 0x67, 0xa7, 0x96, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04,
	0x41, 0x79, 0x4a, 0xaa, 0x5c, 0xdc, 0x60, 0x55, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0xb8, 0x94,
	0x19, 0xf9, 0x72, 0x71, 0x07, 0x81, 0x6d, 0x0b, 0x00, 0x59, 0x26, 0x64, 0xc7, 0xc5, 0x0d, 0x66,
	0x04, 0x80, 0x65, 0x85, 0x44, 0xf5, 0xe0, 0xd6, 0xea, 0x21, 0xec, 0x94, 0x12, 0x43, 0x17, 0x86,
	0x58, 0xa2, 0xc4, 0xe0, 0x24, 0x1d, 0x25, 0x09, 0x96, 0xd2, 0x07, 0xf9, 0x26, 0x39, 0x27, 0xbf,
	0x34, 0x45, 0x3f, 0x3d, 0x1f, 0xea, 0xad, 0x24, 0x36, 0x30, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xd3, 0xb7, 0xe7, 0x39, 0xeb, 0x00, 0x00, 0x00,
}
