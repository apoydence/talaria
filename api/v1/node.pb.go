// Code generated by protoc-gen-go.
// source: node.proto
// DO NOT EDIT!

/*
Package talaria is a generated protocol buffer package.

It is generated from these files:
	node.proto
	scheduler.proto

It has these top-level messages:
	BufferInfo
	WriteDataPacket
	ReadDataPacket
	WriteResponse
	InfoResponse
	ListClustersInfo
	ListClustersResponse
	CreateInfo
	CreateResponse
	ListInfo
	ListResponse
	ClusterInfo
	NodeInfo
*/
package talaria

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

type BufferInfo struct {
	Name         string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	StartIndex   uint64 `protobuf:"varint,2,opt,name=startIndex" json:"startIndex,omitempty"`
	StartFromEnd bool   `protobuf:"varint,3,opt,name=startFromEnd" json:"startFromEnd,omitempty"`
	Tail         bool   `protobuf:"varint,4,opt,name=tail" json:"tail,omitempty"`
}

func (m *BufferInfo) Reset()                    { *m = BufferInfo{} }
func (m *BufferInfo) String() string            { return proto.CompactTextString(m) }
func (*BufferInfo) ProtoMessage()               {}
func (*BufferInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BufferInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BufferInfo) GetStartIndex() uint64 {
	if m != nil {
		return m.StartIndex
	}
	return 0
}

func (m *BufferInfo) GetStartFromEnd() bool {
	if m != nil {
		return m.StartFromEnd
	}
	return false
}

func (m *BufferInfo) GetTail() bool {
	if m != nil {
		return m.Tail
	}
	return false
}

type WriteDataPacket struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Message []byte `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *WriteDataPacket) Reset()                    { *m = WriteDataPacket{} }
func (m *WriteDataPacket) String() string            { return proto.CompactTextString(m) }
func (*WriteDataPacket) ProtoMessage()               {}
func (*WriteDataPacket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WriteDataPacket) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WriteDataPacket) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type ReadDataPacket struct {
	Message []byte `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Index   uint64 `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
}

func (m *ReadDataPacket) Reset()                    { *m = ReadDataPacket{} }
func (m *ReadDataPacket) String() string            { return proto.CompactTextString(m) }
func (*ReadDataPacket) ProtoMessage()               {}
func (*ReadDataPacket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadDataPacket) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *ReadDataPacket) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type WriteResponse struct {
	LastWriteIndex uint64 `protobuf:"varint,1,opt,name=lastWriteIndex" json:"lastWriteIndex,omitempty"`
	Error          string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *WriteResponse) Reset()                    { *m = WriteResponse{} }
func (m *WriteResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()               {}
func (*WriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WriteResponse) GetLastWriteIndex() uint64 {
	if m != nil {
		return m.LastWriteIndex
	}
	return 0
}

func (m *WriteResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type InfoResponse struct {
	Uri string `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
}

func (m *InfoResponse) Reset()                    { *m = InfoResponse{} }
func (m *InfoResponse) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()               {}
func (*InfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *InfoResponse) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type ListClustersInfo struct {
}

func (m *ListClustersInfo) Reset()                    { *m = ListClustersInfo{} }
func (m *ListClustersInfo) String() string            { return proto.CompactTextString(m) }
func (*ListClustersInfo) ProtoMessage()               {}
func (*ListClustersInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ListClustersResponse struct {
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *ListClustersResponse) Reset()                    { *m = ListClustersResponse{} }
func (m *ListClustersResponse) String() string            { return proto.CompactTextString(m) }
func (*ListClustersResponse) ProtoMessage()               {}
func (*ListClustersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListClustersResponse) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func init() {
	proto.RegisterType((*BufferInfo)(nil), "talaria.BufferInfo")
	proto.RegisterType((*WriteDataPacket)(nil), "talaria.WriteDataPacket")
	proto.RegisterType((*ReadDataPacket)(nil), "talaria.ReadDataPacket")
	proto.RegisterType((*WriteResponse)(nil), "talaria.WriteResponse")
	proto.RegisterType((*InfoResponse)(nil), "talaria.InfoResponse")
	proto.RegisterType((*ListClustersInfo)(nil), "talaria.ListClustersInfo")
	proto.RegisterType((*ListClustersResponse)(nil), "talaria.ListClustersResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Node service

type NodeClient interface {
	Write(ctx context.Context, opts ...grpc.CallOption) (Node_WriteClient, error)
	Read(ctx context.Context, in *BufferInfo, opts ...grpc.CallOption) (Node_ReadClient, error)
	ListClusters(ctx context.Context, in *ListClustersInfo, opts ...grpc.CallOption) (*ListClustersResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Write(ctx context.Context, opts ...grpc.CallOption) (Node_WriteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Node_serviceDesc.Streams[0], c.cc, "/talaria.Node/Write", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeWriteClient{stream}
	return x, nil
}

type Node_WriteClient interface {
	Send(*WriteDataPacket) error
	CloseAndRecv() (*WriteResponse, error)
	grpc.ClientStream
}

type nodeWriteClient struct {
	grpc.ClientStream
}

func (x *nodeWriteClient) Send(m *WriteDataPacket) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeWriteClient) CloseAndRecv() (*WriteResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WriteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) Read(ctx context.Context, in *BufferInfo, opts ...grpc.CallOption) (Node_ReadClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Node_serviceDesc.Streams[1], c.cc, "/talaria.Node/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Node_ReadClient interface {
	Recv() (*ReadDataPacket, error)
	grpc.ClientStream
}

type nodeReadClient struct {
	grpc.ClientStream
}

func (x *nodeReadClient) Recv() (*ReadDataPacket, error) {
	m := new(ReadDataPacket)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) ListClusters(ctx context.Context, in *ListClustersInfo, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	out := new(ListClustersResponse)
	err := grpc.Invoke(ctx, "/talaria.Node/ListClusters", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	Write(Node_WriteServer) error
	Read(*BufferInfo, Node_ReadServer) error
	ListClusters(context.Context, *ListClustersInfo) (*ListClustersResponse, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Write_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).Write(&nodeWriteServer{stream})
}

type Node_WriteServer interface {
	SendAndClose(*WriteResponse) error
	Recv() (*WriteDataPacket, error)
	grpc.ServerStream
}

type nodeWriteServer struct {
	grpc.ServerStream
}

func (x *nodeWriteServer) SendAndClose(m *WriteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeWriteServer) Recv() (*WriteDataPacket, error) {
	m := new(WriteDataPacket)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Node_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BufferInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServer).Read(m, &nodeReadServer{stream})
}

type Node_ReadServer interface {
	Send(*ReadDataPacket) error
	grpc.ServerStream
}

type nodeReadServer struct {
	grpc.ServerStream
}

func (x *nodeReadServer) Send(m *ReadDataPacket) error {
	return x.ServerStream.SendMsg(m)
}

func _Node_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ListClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/talaria.Node/ListClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ListClusters(ctx, req.(*ListClustersInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "talaria.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListClusters",
			Handler:    _Node_ListClusters_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Write",
			Handler:       _Node_Write_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Read",
			Handler:       _Node_Read_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "node.proto",
}

func init() { proto.RegisterFile("node.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0xed, 0xfe, 0x9a, 0xfe, 0x6a, 0x87, 0x58, 0xcb, 0x58, 0x34, 0x16, 0x94, 0xb0, 0x07, 0xc9,
	0x41, 0x8a, 0xe8, 0xc5, 0x8b, 0x28, 0xfe, 0x83, 0x8a, 0x8a, 0xec, 0xc5, 0xf3, 0x6a, 0xa6, 0x12,
	0x4c, 0xb3, 0x65, 0x77, 0x0b, 0xfd, 0x8e, 0x7e, 0x29, 0xd9, 0x6d, 0x9b, 0x26, 0xa5, 0xb7, 0x99,
	0x37, 0x2f, 0x6f, 0xde, 0xbc, 0x2c, 0x40, 0xa1, 0x52, 0x1a, 0x4e, 0xb5, 0xb2, 0x0a, 0xdb, 0x56,
	0xe6, 0x52, 0x67, 0x92, 0xcf, 0x01, 0xee, 0x66, 0xe3, 0x31, 0xe9, 0x51, 0x31, 0x56, 0x88, 0x10,
	0x14, 0x72, 0x42, 0x11, 0x8b, 0x59, 0xd2, 0x11, 0xbe, 0xc6, 0x13, 0x00, 0x63, 0xa5, 0xb6, 0xa3,
	0x22, 0xa5, 0x79, 0xf4, 0x2f, 0x66, 0x49, 0x20, 0x2a, 0x08, 0x72, 0x08, 0x7d, 0xf7, 0xa4, 0xd5,
	0xe4, 0xb1, 0x48, 0xa3, 0x66, 0xcc, 0x92, 0x1d, 0x51, 0xc3, 0x9c, 0xae, 0x95, 0x59, 0x1e, 0x05,
	0x7e, 0xe6, 0x6b, 0x7e, 0x03, 0x7b, 0x1f, 0x3a, 0xb3, 0xf4, 0x20, 0xad, 0x7c, 0x97, 0x5f, 0x3f,
	0x64, 0xb7, 0xae, 0x8f, 0xa0, 0x3d, 0x21, 0x63, 0xe4, 0x37, 0xf9, 0xdd, 0xa1, 0x58, 0xb5, 0xfc,
	0x16, 0xba, 0x82, 0x64, 0x5a, 0xf9, 0xbe, 0xc2, 0x65, 0x35, 0x2e, 0xf6, 0xa1, 0x95, 0x55, 0xfc,
	0x2f, 0x1a, 0xfe, 0x0a, 0xbb, 0xde, 0x82, 0x20, 0x33, 0x55, 0x85, 0x21, 0x3c, 0x85, 0x6e, 0x2e,
	0x8d, 0xf5, 0xe0, 0xe2, 0x5e, 0xe6, 0xf9, 0x1b, 0xa8, 0x93, 0x23, 0xad, 0x95, 0xf6, 0x72, 0x1d,
	0xb1, 0x68, 0x78, 0x0c, 0xa1, 0x4b, 0xb1, 0x54, 0xeb, 0x41, 0x73, 0xa6, 0xb3, 0xe5, 0x35, 0xae,
	0xe4, 0x08, 0xbd, 0x97, 0xcc, 0xd8, 0xfb, 0x7c, 0x66, 0x2c, 0x69, 0xe3, 0xd8, 0xfc, 0x0c, 0xfa,
	0x55, 0xac, 0xfc, 0xba, 0x0f, 0x2d, 0x17, 0x80, 0x89, 0x58, 0xdc, 0x74, 0x3b, 0x7c, 0x73, 0xf1,
	0xcb, 0x20, 0x78, 0x53, 0x29, 0xe1, 0x35, 0xb4, 0xbc, 0x21, 0x8c, 0x86, 0xcb, 0x7f, 0x39, 0xdc,
	0x88, 0x73, 0x70, 0x50, 0x9f, 0xac, 0x94, 0x79, 0x23, 0x61, 0x78, 0x05, 0x81, 0x0b, 0x0f, 0xf7,
	0x4b, 0xce, 0xfa, 0x19, 0x0c, 0x0e, 0x4b, 0xb0, 0x1e, 0x30, 0x6f, 0x9c, 0x33, 0x7c, 0x86, 0xb0,
	0xea, 0x17, 0x8f, 0x4a, 0xf2, 0xe6, 0x69, 0x83, 0xe3, 0xad, 0xa3, 0xb5, 0x8f, 0xcf, 0xff, 0xfe,
	0x35, 0x5e, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x2e, 0xad, 0xba, 0x9b, 0x02, 0x00, 0x00,
}