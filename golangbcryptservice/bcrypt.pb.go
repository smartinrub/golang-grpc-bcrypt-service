// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bcrypt.proto

package bcrypt

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

type PasswordRequest struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswordRequest) Reset()         { *m = PasswordRequest{} }
func (m *PasswordRequest) String() string { return proto.CompactTextString(m) }
func (*PasswordRequest) ProtoMessage()    {}
func (*PasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcrypt_f7f3b704c4ab7630, []int{0}
}
func (m *PasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordRequest.Unmarshal(m, b)
}
func (m *PasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordRequest.Marshal(b, m, deterministic)
}
func (dst *PasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordRequest.Merge(dst, src)
}
func (m *PasswordRequest) XXX_Size() int {
	return xxx_messageInfo_PasswordRequest.Size(m)
}
func (m *PasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordRequest proto.InternalMessageInfo

func (m *PasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type PasswordHashedResponse struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswordHashedResponse) Reset()         { *m = PasswordHashedResponse{} }
func (m *PasswordHashedResponse) String() string { return proto.CompactTextString(m) }
func (*PasswordHashedResponse) ProtoMessage()    {}
func (*PasswordHashedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcrypt_f7f3b704c4ab7630, []int{1}
}
func (m *PasswordHashedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordHashedResponse.Unmarshal(m, b)
}
func (m *PasswordHashedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordHashedResponse.Marshal(b, m, deterministic)
}
func (dst *PasswordHashedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordHashedResponse.Merge(dst, src)
}
func (m *PasswordHashedResponse) XXX_Size() int {
	return xxx_messageInfo_PasswordHashedResponse.Size(m)
}
func (m *PasswordHashedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordHashedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordHashedResponse proto.InternalMessageInfo

func (m *PasswordHashedResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PasswordRequest)(nil), "PasswordRequest")
	proto.RegisterType((*PasswordHashedResponse)(nil), "PasswordHashedResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BCryptClient is the client API for BCrypt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BCryptClient interface {
	PostBCrypt(ctx context.Context, in *PasswordRequest, opts ...grpc.CallOption) (*PasswordHashedResponse, error)
}

type bCryptClient struct {
	cc *grpc.ClientConn
}

func NewBCryptClient(cc *grpc.ClientConn) BCryptClient {
	return &bCryptClient{cc}
}

func (c *bCryptClient) PostBCrypt(ctx context.Context, in *PasswordRequest, opts ...grpc.CallOption) (*PasswordHashedResponse, error) {
	out := new(PasswordHashedResponse)
	err := c.cc.Invoke(ctx, "/BCrypt/PostBCrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BCryptServer is the server API for BCrypt service.
type BCryptServer interface {
	PostBCrypt(context.Context, *PasswordRequest) (*PasswordHashedResponse, error)
}

func RegisterBCryptServer(s *grpc.Server, srv BCryptServer) {
	s.RegisterService(&_BCrypt_serviceDesc, srv)
}

func _BCrypt_PostBCrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BCryptServer).PostBCrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BCrypt/PostBCrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BCryptServer).PostBCrypt(ctx, req.(*PasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BCrypt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BCrypt",
	HandlerType: (*BCryptServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostBCrypt",
			Handler:    _BCrypt_PostBCrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bcrypt.proto",
}

func init() { proto.RegisterFile("bcrypt.proto", fileDescriptor_bcrypt_f7f3b704c4ab7630) }

var fileDescriptor_bcrypt_f7f3b704c4ab7630 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x4a, 0x2e, 0xaa,
	0x2c, 0x28, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe5, 0xe2, 0x0f, 0x48, 0x2c, 0x2e,
	0x2e, 0xcf, 0x2f, 0x4a, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0x28,
	0x80, 0x0a, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x4a, 0x3a, 0x5c, 0x62, 0x30,
	0xe5, 0x1e, 0x89, 0xc5, 0x19, 0xa9, 0x29, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42,
	0x42, 0x5c, 0x2c, 0x19, 0x89, 0xc5, 0x19, 0x50, 0x1d, 0x60, 0xb6, 0x91, 0x33, 0x17, 0x9b, 0x93,
	0x33, 0xc8, 0x32, 0x21, 0x4b, 0x2e, 0xae, 0x80, 0xfc, 0xe2, 0x12, 0x28, 0x4f, 0x40, 0x0f, 0xcd,
	0x4e, 0x29, 0x71, 0x3d, 0xec, 0xc6, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0x1d, 0x6a, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x13, 0xfd, 0xbe, 0x2f, 0xb8, 0x00, 0x00, 0x00,
}