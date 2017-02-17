// Code generated by protoc-gen-go.
// source: scheduler.proto
// DO NOT EDIT!

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

type CreateInfo struct {
	Name       string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	BufferSize uint64 `protobuf:"varint,2,opt,name=buffer_size,json=bufferSize" json:"buffer_size,omitempty"`
}

func (m *CreateInfo) Reset()                    { *m = CreateInfo{} }
func (m *CreateInfo) String() string            { return proto.CompactTextString(m) }
func (*CreateInfo) ProtoMessage()               {}
func (*CreateInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateInfo) GetBufferSize() uint64 {
	if m != nil {
		return m.BufferSize
	}
	return 0
}

type CreateResponse struct {
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type ListInfo struct {
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *ListInfo) Reset()                    { *m = ListInfo{} }
func (m *ListInfo) String() string            { return proto.CompactTextString(m) }
func (*ListInfo) ProtoMessage()               {}
func (*ListInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ListInfo) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type ListResponse struct {
	Info []*ClusterInfo `protobuf:"bytes,1,rep,name=info" json:"info,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ListResponse) GetInfo() []*ClusterInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ClusterInfo struct {
	Name   string      `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Leader string      `protobuf:"bytes,2,opt,name=leader" json:"leader,omitempty"`
	Nodes  []*NodeInfo `protobuf:"bytes,3,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *ClusterInfo) Reset()                    { *m = ClusterInfo{} }
func (m *ClusterInfo) String() string            { return proto.CompactTextString(m) }
func (*ClusterInfo) ProtoMessage()               {}
func (*ClusterInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ClusterInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterInfo) GetLeader() string {
	if m != nil {
		return m.Leader
	}
	return ""
}

func (m *ClusterInfo) GetNodes() []*NodeInfo {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type NodeInfo struct {
	URI string `protobuf:"bytes,1,opt,name=URI,json=uRI" json:"URI,omitempty"`
}

func (m *NodeInfo) Reset()                    { *m = NodeInfo{} }
func (m *NodeInfo) String() string            { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()               {}
func (*NodeInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *NodeInfo) GetURI() string {
	if m != nil {
		return m.URI
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateInfo)(nil), "talaria.CreateInfo")
	proto.RegisterType((*CreateResponse)(nil), "talaria.CreateResponse")
	proto.RegisterType((*ListInfo)(nil), "talaria.ListInfo")
	proto.RegisterType((*ListResponse)(nil), "talaria.ListResponse")
	proto.RegisterType((*ClusterInfo)(nil), "talaria.ClusterInfo")
	proto.RegisterType((*NodeInfo)(nil), "talaria.NodeInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Scheduler service

type SchedulerClient interface {
	Create(ctx context.Context, in *CreateInfo, opts ...grpc.CallOption) (*CreateResponse, error)
	ListClusterInfo(ctx context.Context, in *ListInfo, opts ...grpc.CallOption) (*ListResponse, error)
}

type schedulerClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerClient(cc *grpc.ClientConn) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) Create(ctx context.Context, in *CreateInfo, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/talaria.Scheduler/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) ListClusterInfo(ctx context.Context, in *ListInfo, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/talaria.Scheduler/ListClusterInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Scheduler service

type SchedulerServer interface {
	Create(context.Context, *CreateInfo) (*CreateResponse, error)
	ListClusterInfo(context.Context, *ListInfo) (*ListResponse, error)
}

func RegisterSchedulerServer(s *grpc.Server, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/talaria.Scheduler/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).Create(ctx, req.(*CreateInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_ListClusterInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).ListClusterInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/talaria.Scheduler/ListClusterInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).ListClusterInfo(ctx, req.(*ListInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "talaria.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Scheduler_Create_Handler,
		},
		{
			MethodName: "ListClusterInfo",
			Handler:    _Scheduler_ListClusterInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheduler.proto",
}

func init() { proto.RegisterFile("scheduler.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x4f, 0x83, 0x40,
	0x10, 0x15, 0x69, 0xb1, 0x0c, 0xc6, 0xd6, 0xb1, 0x2a, 0x69, 0x4c, 0x24, 0x7b, 0x91, 0x13, 0x87,
	0x7a, 0xe9, 0xc5, 0x83, 0xe9, 0x89, 0xc4, 0x78, 0xd8, 0xc6, 0xb3, 0x59, 0x64, 0x88, 0x24, 0xc8,
	0x36, 0xbb, 0x70, 0xe9, 0xd9, 0x1f, 0x6e, 0x58, 0xbe, 0xaa, 0xf1, 0xc6, 0x3c, 0xde, 0xbc, 0x79,
	0xef, 0x2d, 0xcc, 0xf5, 0xc7, 0x27, 0xa5, 0x75, 0x41, 0x2a, 0xda, 0x2b, 0x59, 0x49, 0x3c, 0xab,
	0x44, 0x21, 0x54, 0x2e, 0xd8, 0x33, 0xc0, 0x56, 0x91, 0xa8, 0x28, 0x2e, 0x33, 0x89, 0x08, 0x93,
	0x52, 0x7c, 0x91, 0x6f, 0x05, 0x56, 0xe8, 0x72, 0xf3, 0x8d, 0xf7, 0xe0, 0x25, 0x75, 0x96, 0x91,
	0x7a, 0xd7, 0xf9, 0x81, 0xfc, 0xd3, 0xc0, 0x0a, 0x27, 0x1c, 0x5a, 0x68, 0x97, 0x1f, 0x88, 0x2d,
	0xe0, 0xa2, 0x95, 0xe0, 0xa4, 0xf7, 0xb2, 0xd4, 0xc4, 0x02, 0x98, 0xbd, 0xe4, 0xba, 0x32, 0x92,
	0x4b, 0x98, 0x36, 0x32, 0xda, 0xb7, 0x02, 0x3b, 0x74, 0x79, 0x3b, 0xb0, 0x0d, 0x9c, 0x37, 0x8c,
	0x7e, 0x03, 0x43, 0x98, 0xe4, 0x65, 0x26, 0x0d, 0xc9, 0x5b, 0x2f, 0xa3, 0xce, 0x5e, 0xb4, 0x2d,
	0x6a, 0x5d, 0x91, 0x6a, 0x94, 0xb8, 0x61, 0xb0, 0x04, 0xbc, 0x23, 0xf0, 0x5f, 0xc7, 0x37, 0xe0,
	0x14, 0x24, 0x52, 0x52, 0xc6, 0xac, 0xcb, 0xbb, 0x09, 0x1f, 0x60, 0x5a, 0xca, 0x94, 0xb4, 0x6f,
	0x9b, 0x2b, 0x97, 0xc3, 0x95, 0x57, 0x99, 0x9a, 0xfc, 0xbc, 0xfd, 0xcf, 0xee, 0x60, 0xd6, 0x43,
	0xb8, 0x00, 0xfb, 0x8d, 0xc7, 0x9d, 0xbe, 0x5d, 0xf3, 0x78, 0xfd, 0x6d, 0x81, 0xbb, 0xeb, 0xfb,
	0xc4, 0x0d, 0x38, 0x6d, 0x7a, 0xbc, 0x1a, 0x5d, 0x0f, 0x8d, 0xae, 0x6e, 0xff, 0x80, 0x43, 0x47,
	0x27, 0xf8, 0x04, 0xf3, 0xa6, 0x83, 0xe3, 0x34, 0xa3, 0xa5, 0xbe, 0xbf, 0xd5, 0xf5, 0x2f, 0x68,
	0x5c, 0x4f, 0x1c, 0xf3, 0x92, 0x8f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x64, 0x4e, 0x97,
	0xdc, 0x01, 0x00, 0x00,
}
