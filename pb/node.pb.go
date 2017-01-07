// Code generated by protoc-gen-go.
// source: node.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	node.proto
	scheduler.proto

It has these top-level messages:
	BufferInfo
	WriteDataPacket
	ReadDataPacket
	WriteResponse
	InfoResponse
	CreateInfo
	CreateResponse
	ListInfo
	ListResponse
	ClusterInfo
	NodeInfo
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

type BufferInfo struct {
	Name         string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	StartIndex   uint64 `protobuf:"varint,2,opt,name=startIndex" json:"startIndex,omitempty"`
	StartFromEnd bool   `protobuf:"varint,3,opt,name=startFromEnd" json:"startFromEnd,omitempty"`
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

func init() {
	proto.RegisterType((*BufferInfo)(nil), "pb.BufferInfo")
	proto.RegisterType((*WriteDataPacket)(nil), "pb.WriteDataPacket")
	proto.RegisterType((*ReadDataPacket)(nil), "pb.ReadDataPacket")
	proto.RegisterType((*WriteResponse)(nil), "pb.WriteResponse")
	proto.RegisterType((*InfoResponse)(nil), "pb.InfoResponse")
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
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Write(ctx context.Context, opts ...grpc.CallOption) (Node_WriteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Node_serviceDesc.Streams[0], c.cc, "/pb.Node/Write", opts...)
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
	stream, err := grpc.NewClientStream(ctx, &_Node_serviceDesc.Streams[1], c.cc, "/pb.Node/Read", opts...)
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

// Server API for Node service

type NodeServer interface {
	Write(Node_WriteServer) error
	Read(*BufferInfo, Node_ReadServer) error
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

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Node",
	HandlerType: (*NodeServer)(nil),
	Methods:     []grpc.MethodDesc{},
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
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x51, 0xc1, 0x4e, 0xeb, 0x30,
	0x10, 0xac, 0xdb, 0xf4, 0x3d, 0xba, 0x2a, 0x05, 0x16, 0x0e, 0x51, 0x0f, 0x28, 0xf2, 0x01, 0xe5,
	0x14, 0x21, 0xfa, 0x01, 0x20, 0x04, 0x48, 0x3d, 0x80, 0x90, 0x2f, 0x9c, 0x1d, 0xbc, 0x41, 0x51,
	0x89, 0x1d, 0xd9, 0xae, 0xc4, 0xe7, 0x23, 0x3b, 0xb4, 0x24, 0x15, 0xb7, 0x9d, 0xf1, 0x7a, 0x76,
	0x67, 0x16, 0x40, 0x1b, 0x45, 0x45, 0x6b, 0x8d, 0x37, 0x38, 0x6e, 0x4b, 0xae, 0x00, 0xee, 0xb7,
	0x55, 0x45, 0x76, 0xad, 0x2b, 0x83, 0x08, 0x89, 0x96, 0x0d, 0xa5, 0x2c, 0x63, 0xf9, 0x4c, 0xc4,
	0x1a, 0x2f, 0x01, 0x9c, 0x97, 0xd6, 0xaf, 0xb5, 0xa2, 0xaf, 0x74, 0x9c, 0xb1, 0x3c, 0x11, 0x3d,
	0x06, 0x39, 0xcc, 0x23, 0x7a, 0xb2, 0xa6, 0x79, 0xd4, 0x2a, 0x9d, 0x64, 0x2c, 0x3f, 0x12, 0x03,
	0x8e, 0xdf, 0xc2, 0xc9, 0x9b, 0xad, 0x3d, 0x3d, 0x48, 0x2f, 0x5f, 0xe5, 0xfb, 0x86, 0xfc, 0x9f,
	0xa3, 0x52, 0xf8, 0xdf, 0x90, 0x73, 0xf2, 0x83, 0xe2, 0x9c, 0xb9, 0xd8, 0x41, 0x7e, 0x07, 0x0b,
	0x41, 0x52, 0xf5, 0xfe, 0xf7, 0x7a, 0xd9, 0xa0, 0x17, 0x2f, 0x60, 0x5a, 0xf7, 0x76, 0xed, 0x00,
	0x7f, 0x86, 0xe3, 0xb8, 0x82, 0x20, 0xd7, 0x1a, 0xed, 0x08, 0xaf, 0x60, 0xf1, 0x29, 0x9d, 0x8f,
	0x64, 0xe7, 0x8d, 0xc5, 0xfe, 0x03, 0x36, 0xc8, 0x91, 0xb5, 0xc6, 0x46, 0xb9, 0x99, 0xe8, 0x00,
	0xcf, 0x60, 0x1e, 0x12, 0xdb, 0xab, 0x9d, 0xc2, 0x64, 0x6b, 0xeb, 0x1f, 0x37, 0xa1, 0xbc, 0xd9,
	0x40, 0xf2, 0x62, 0x14, 0xe1, 0x0a, 0xa6, 0x51, 0x0d, 0xcf, 0x8b, 0xb6, 0x2c, 0x0e, 0x62, 0x58,
	0x9e, 0xed, 0xc9, 0x9d, 0x14, 0x1f, 0xe5, 0x0c, 0x0b, 0x48, 0x82, 0x5f, 0x5c, 0x84, 0xe7, 0xdf,
	0x03, 0x2d, 0x31, 0xe0, 0x61, 0x12, 0x7c, 0x74, 0xcd, 0xca, 0x7f, 0xf1, 0xa2, 0xab, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1a, 0xe9, 0x4d, 0x0c, 0xdf, 0x01, 0x00, 0x00,
}