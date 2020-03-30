// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sum.proto

package sum

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

type SumRequest struct {
	Input                string   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62743f9cdc99b9fd, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

type SumResponse struct {
	Output               string   `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62743f9cdc99b9fd, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "SumRequest")
	proto.RegisterType((*SumResponse)(nil), "SumResponse")
}

func init() { proto.RegisterFile("sum.proto", fileDescriptor_62743f9cdc99b9fd) }

var fileDescriptor_62743f9cdc99b9fd = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2e, 0xcd, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe2, 0xe2, 0x0a, 0x2e, 0xcd, 0x0d, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0xcd, 0xcc, 0x2b, 0x28, 0x2d, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0x54, 0xb9, 0xb8, 0xc1, 0x6a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a,
	0x53, 0x85, 0xc4, 0xb8, 0xd8, 0xf2, 0x4b, 0x4b, 0x10, 0xaa, 0xa0, 0x3c, 0x23, 0x43, 0xb0, 0x51,
	0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xca, 0x5c, 0x6c, 0xe9, 0xa9, 0x25, 0xc1, 0xa5,
	0xb9, 0x42, 0xdc, 0x7a, 0x08, 0x1b, 0xa4, 0x78, 0xf4, 0x90, 0x8c, 0x4a, 0x62, 0x03, 0x3b, 0xc2,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x14, 0xf6, 0x9e, 0x91, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumServiceClient interface {
	GetSum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type sumServiceClient struct {
	cc *grpc.ClientConn
}

func NewSumServiceClient(cc *grpc.ClientConn) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) GetSum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/SumService/getSum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SumServiceServer is the server API for SumService service.
type SumServiceServer interface {
	GetSum(context.Context, *SumRequest) (*SumResponse, error)
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_GetSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).GetSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SumService/GetSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).GetSum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getSum",
			Handler:    _SumService_GetSum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sum.proto",
}