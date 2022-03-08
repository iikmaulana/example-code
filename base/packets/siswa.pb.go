// Code generated by protoc-gen-go. DO NOT EDIT.
// source: siswa.proto

package packets

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SiswaRequest struct {
	Data                 *any.Any `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SiswaRequest) Reset()         { *m = SiswaRequest{} }
func (m *SiswaRequest) String() string { return proto.CompactTextString(m) }
func (*SiswaRequest) ProtoMessage()    {}
func (*SiswaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9a300d1b80ca330, []int{0}
}

func (m *SiswaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SiswaRequest.Unmarshal(m, b)
}
func (m *SiswaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SiswaRequest.Marshal(b, m, deterministic)
}
func (m *SiswaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SiswaRequest.Merge(m, src)
}
func (m *SiswaRequest) XXX_Size() int {
	return xxx_messageInfo_SiswaRequest.Size(m)
}
func (m *SiswaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SiswaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SiswaRequest proto.InternalMessageInfo

func (m *SiswaRequest) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type OutputSiswa struct {
	Status               int32    `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Data                 *any.Any `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutputSiswa) Reset()         { *m = OutputSiswa{} }
func (m *OutputSiswa) String() string { return proto.CompactTextString(m) }
func (*OutputSiswa) ProtoMessage()    {}
func (*OutputSiswa) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9a300d1b80ca330, []int{1}
}

func (m *OutputSiswa) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputSiswa.Unmarshal(m, b)
}
func (m *OutputSiswa) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputSiswa.Marshal(b, m, deterministic)
}
func (m *OutputSiswa) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputSiswa.Merge(m, src)
}
func (m *OutputSiswa) XXX_Size() int {
	return xxx_messageInfo_OutputSiswa.Size(m)
}
func (m *OutputSiswa) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputSiswa.DiscardUnknown(m)
}

var xxx_messageInfo_OutputSiswa proto.InternalMessageInfo

func (m *OutputSiswa) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *OutputSiswa) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SiswaRequest)(nil), "packets.SiswaRequest")
	proto.RegisterType((*OutputSiswa)(nil), "packets.OutputSiswa")
}

func init() { proto.RegisterFile("siswa.proto", fileDescriptor_e9a300d1b80ca330) }

var fileDescriptor_e9a300d1b80ca330 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xce, 0x2c, 0x2e,
	0x4f, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0x29,
	0x96, 0x92, 0x4c, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x27, 0x95, 0xa6, 0xe9, 0x27,
	0xe6, 0x55, 0x42, 0xd4, 0x28, 0x59, 0x70, 0xf1, 0x04, 0x83, 0xb4, 0x04, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x08, 0x69, 0x70, 0xb1, 0xb8, 0x24, 0x96, 0x24, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x1b, 0x89, 0xe8, 0x41, 0x74, 0xea, 0xc1, 0x74, 0xea, 0x39, 0xe6, 0x55, 0x06, 0x81, 0x55, 0x28,
	0xf9, 0x73, 0x71, 0xfb, 0x97, 0x96, 0x14, 0x94, 0x96, 0x80, 0xf5, 0x0b, 0x89, 0x71, 0xb1, 0x05,
	0x97, 0x24, 0x96, 0x94, 0x16, 0x83, 0xb5, 0xb2, 0x06, 0x41, 0x79, 0x70, 0x03, 0x99, 0x08, 0x19,
	0x68, 0x34, 0x89, 0x91, 0x8b, 0x15, 0x62, 0x96, 0x0b, 0x97, 0xb0, 0x7b, 0x6a, 0x89, 0x4f, 0x66,
	0x31, 0xc4, 0xec, 0xd0, 0xe2, 0xd4, 0xe4, 0xc4, 0xe2, 0x54, 0x21, 0x51, 0x3d, 0xa8, 0x87, 0xf4,
	0x90, 0x9d, 0x2c, 0x25, 0x02, 0x17, 0x46, 0x72, 0x8f, 0x12, 0x83, 0x90, 0x33, 0x97, 0x90, 0x73,
	0x51, 0x6a, 0x62, 0x49, 0x2a, 0x05, 0x86, 0x38, 0x71, 0x47, 0x71, 0xea, 0x59, 0x43, 0xa5, 0x92,
	0xd8, 0xc0, 0xae, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xff, 0x22, 0xd1, 0x58, 0x66, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SiswaClient is the client API for Siswa service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SiswaClient interface {
	GetListSiswaUsecase(ctx context.Context, in *SiswaRequest, opts ...grpc.CallOption) (*OutputSiswa, error)
	CreateSiswaUsecase(ctx context.Context, in *SiswaRequest, opts ...grpc.CallOption) (*OutputSiswa, error)
}

type siswaClient struct {
	cc *grpc.ClientConn
}

func NewSiswaClient(cc *grpc.ClientConn) SiswaClient {
	return &siswaClient{cc}
}

func (c *siswaClient) GetListSiswaUsecase(ctx context.Context, in *SiswaRequest, opts ...grpc.CallOption) (*OutputSiswa, error) {
	out := new(OutputSiswa)
	err := c.cc.Invoke(ctx, "/packets.Siswa/GetListSiswaUsecase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siswaClient) CreateSiswaUsecase(ctx context.Context, in *SiswaRequest, opts ...grpc.CallOption) (*OutputSiswa, error) {
	out := new(OutputSiswa)
	err := c.cc.Invoke(ctx, "/packets.Siswa/CreateSiswaUsecase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SiswaServer is the server API for Siswa service.
type SiswaServer interface {
	GetListSiswaUsecase(context.Context, *SiswaRequest) (*OutputSiswa, error)
	CreateSiswaUsecase(context.Context, *SiswaRequest) (*OutputSiswa, error)
}

// UnimplementedSiswaServer can be embedded to have forward compatible implementations.
type UnimplementedSiswaServer struct {
}

func (*UnimplementedSiswaServer) GetListSiswaUsecase(ctx context.Context, req *SiswaRequest) (*OutputSiswa, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListSiswaUsecase not implemented")
}
func (*UnimplementedSiswaServer) CreateSiswaUsecase(ctx context.Context, req *SiswaRequest) (*OutputSiswa, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSiswaUsecase not implemented")
}

func RegisterSiswaServer(s *grpc.Server, srv SiswaServer) {
	s.RegisterService(&_Siswa_serviceDesc, srv)
}

func _Siswa_GetListSiswaUsecase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SiswaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiswaServer).GetListSiswaUsecase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.Siswa/GetListSiswaUsecase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiswaServer).GetListSiswaUsecase(ctx, req.(*SiswaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Siswa_CreateSiswaUsecase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SiswaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiswaServer).CreateSiswaUsecase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/packets.Siswa/CreateSiswaUsecase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiswaServer).CreateSiswaUsecase(ctx, req.(*SiswaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Siswa_serviceDesc = grpc.ServiceDesc{
	ServiceName: "packets.Siswa",
	HandlerType: (*SiswaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetListSiswaUsecase",
			Handler:    _Siswa_GetListSiswaUsecase_Handler,
		},
		{
			MethodName: "CreateSiswaUsecase",
			Handler:    _Siswa_CreateSiswaUsecase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "siswa.proto",
}
