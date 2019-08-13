// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fabex.proto

package fabex

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Request struct {
	Startblock           int64    `protobuf:"varint,1,opt,name=startblock,proto3" json:"startblock,omitempty"`
	Endblock             int64    `protobuf:"varint,2,opt,name=endblock,proto3" json:"endblock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7d7206373264ff4, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetStartblock() int64 {
	if m != nil {
		return m.Startblock
	}
	return 0
}

func (m *Request) GetEndblock() int64 {
	if m != nil {
		return m.Endblock
	}
	return 0
}

type Reply struct {
	Txid                 string   `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Blocknum             int64    `protobuf:"varint,3,opt,name=blocknum,proto3" json:"blocknum,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7d7206373264ff4, []int{1}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *Reply) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Reply) GetBlocknum() int64 {
	if m != nil {
		return m.Blocknum
	}
	return 0
}

func (m *Reply) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "fabex.Request")
	proto.RegisterType((*Reply)(nil), "fabex.Reply")
}

func init() { proto.RegisterFile("fabex.proto", fileDescriptor_d7d7206373264ff4) }

var fileDescriptor_d7d7206373264ff4 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xc1, 0x8a, 0x83, 0x40,
	0x10, 0x44, 0x77, 0x56, 0x5d, 0xd7, 0x5e, 0xd9, 0x43, 0x9f, 0x06, 0x0f, 0x8b, 0x78, 0x12, 0x16,
	0x24, 0x24, 0xf9, 0x05, 0xf3, 0x01, 0xf3, 0x07, 0x63, 0x9c, 0x60, 0xc8, 0xc4, 0x99, 0xe8, 0x08,
	0xfa, 0xf7, 0xc1, 0x36, 0x4a, 0x6e, 0xfd, 0xaa, 0xa8, 0xa2, 0x0b, 0x7e, 0x2e, 0xb2, 0x52, 0x63,
	0x61, 0x3b, 0xe3, 0x0c, 0x06, 0x04, 0x59, 0x09, 0xa1, 0x50, 0x8f, 0x41, 0xf5, 0x0e, 0xff, 0x00,
	0x7a, 0x27, 0x3b, 0x57, 0x69, 0x73, 0xbe, 0x71, 0x96, 0xb2, 0xdc, 0x13, 0x6f, 0x0a, 0x26, 0xf0,
	0xad, 0xda, 0x7a, 0x71, 0x3f, 0xc9, 0xdd, 0x38, 0x53, 0x10, 0x08, 0x65, 0xf5, 0x84, 0x08, 0xbe,
	0x1b, 0xaf, 0x35, 0xc5, 0x23, 0x41, 0xf7, 0xac, 0x35, 0xb2, 0x6f, 0x28, 0x14, 0x09, 0xba, 0xe7,
	0x32, 0x4a, 0xb6, 0xc3, 0x9d, 0x7b, 0x4b, 0xd9, 0xca, 0xc8, 0x21, 0xb4, 0x72, 0xd2, 0x46, 0xd6,
	0xdc, 0x4f, 0x59, 0x1e, 0x8b, 0x15, 0xf7, 0x47, 0x08, 0x4e, 0xf3, 0xdb, 0xf8, 0x0f, 0x61, 0x39,
	0x5a, 0x6d, 0x3a, 0x85, 0xbf, 0xc5, 0x32, 0xeb, 0x35, 0x23, 0x89, 0x37, 0xb6, 0x7a, 0xca, 0x3e,
	0x76, 0xac, 0xfa, 0xa2, 0xc5, 0x87, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x60, 0x8f, 0x30, 0x5d,
	0x00, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FabexClient is the client API for Fabex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FabexClient interface {
	Explore(ctx context.Context, in *Request, opts ...grpc.CallOption) (Fabex_ExploreClient, error)
}

type fabexClient struct {
	cc *grpc.ClientConn
}

func NewFabexClient(cc *grpc.ClientConn) FabexClient {
	return &fabexClient{cc}
}

func (c *fabexClient) Explore(ctx context.Context, in *Request, opts ...grpc.CallOption) (Fabex_ExploreClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fabex_serviceDesc.Streams[0], "/fabex.Fabex/Explore", opts...)
	if err != nil {
		return nil, err
	}
	x := &fabexExploreClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fabex_ExploreClient interface {
	Recv() (*Reply, error)
	grpc.ClientStream
}

type fabexExploreClient struct {
	grpc.ClientStream
}

func (x *fabexExploreClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FabexServer is the server API for Fabex service.
type FabexServer interface {
	Explore(*Request, Fabex_ExploreServer) error
}

// UnimplementedFabexServer can be embedded to have forward compatible implementations.
type UnimplementedFabexServer struct {
}

func (*UnimplementedFabexServer) Explore(req *Request, srv Fabex_ExploreServer) error {
	return status.Errorf(codes.Unimplemented, "method Explore not implemented")
}

func RegisterFabexServer(s *grpc.Server, srv FabexServer) {
	s.RegisterService(&_Fabex_serviceDesc, srv)
}

func _Fabex_Explore_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FabexServer).Explore(m, &fabexExploreServer{stream})
}

type Fabex_ExploreServer interface {
	Send(*Reply) error
	grpc.ServerStream
}

type fabexExploreServer struct {
	grpc.ServerStream
}

func (x *fabexExploreServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

var _Fabex_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fabex.Fabex",
	HandlerType: (*FabexServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Explore",
			Handler:       _Fabex_Explore_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "fabex.proto",
}
