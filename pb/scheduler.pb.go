// Code generated by protoc-gen-go.
// source: scheduler.proto
// DO NOT EDIT!

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

type CreateInfo struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CreateInfo) Reset()                    { *m = CreateInfo{} }
func (m *CreateInfo) String() string            { return proto.CompactTextString(m) }
func (*CreateInfo) ProtoMessage()               {}
func (*CreateInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

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

func init() {
	proto.RegisterType((*CreateInfo)(nil), "pb.CreateInfo")
	proto.RegisterType((*CreateResponse)(nil), "pb.CreateResponse")
	proto.RegisterType((*ListInfo)(nil), "pb.ListInfo")
	proto.RegisterType((*ListResponse)(nil), "pb.ListResponse")
	proto.RegisterType((*ClusterInfo)(nil), "pb.ClusterInfo")
	proto.RegisterType((*NodeInfo)(nil), "pb.NodeInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

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
	err := grpc.Invoke(ctx, "/pb.Scheduler/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) ListClusterInfo(ctx context.Context, in *ListInfo, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/pb.Scheduler/ListClusterInfo", in, out, c.cc, opts...)
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
		FullMethod: "/pb.Scheduler/Create",
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
		FullMethod: "/pb.Scheduler/ListClusterInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).ListClusterInfo(ctx, req.(*ListInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Scheduler",
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
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("scheduler.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x51, 0x31, 0x6f, 0xb3, 0x30,
	0x10, 0xfd, 0x08, 0x09, 0x0a, 0x97, 0x28, 0xa0, 0xd3, 0xa7, 0x0a, 0x45, 0x1d, 0x90, 0xbb, 0x64,
	0x62, 0x08, 0x3f, 0x21, 0x13, 0x52, 0xd5, 0xc1, 0x55, 0xc7, 0x0e, 0x50, 0x2e, 0x6a, 0x24, 0x6a,
	0x5b, 0x36, 0xfc, 0xff, 0xca, 0x47, 0x9c, 0x64, 0xe8, 0xe6, 0x77, 0xf7, 0xee, 0xee, 0xbd, 0x67,
	0xc8, 0xdc, 0xd7, 0x37, 0xf5, 0xd3, 0x40, 0xb6, 0x32, 0x56, 0x8f, 0x1a, 0x17, 0xa6, 0x13, 0x25,
	0xc0, 0xc9, 0x52, 0x3b, 0x52, 0xa3, 0xce, 0x1a, 0x11, 0x96, 0xaa, 0xfd, 0xa1, 0x22, 0x2a, 0xa3,
	0x43, 0x2a, 0xf9, 0x2d, 0x72, 0xd8, 0xcd, 0x0c, 0x49, 0xce, 0x68, 0xe5, 0x48, 0x94, 0xb0, 0x7e,
	0xbd, 0xb8, 0x91, 0x27, 0xfe, 0xc3, 0xca, 0xb3, 0x5c, 0x11, 0x95, 0xf1, 0x21, 0x95, 0x33, 0x10,
	0x35, 0x6c, 0x3d, 0x23, 0x4c, 0xe0, 0x0b, 0x2c, 0x2f, 0xea, 0xac, 0x99, 0xb4, 0x39, 0x66, 0x95,
	0xe9, 0xaa, 0xd3, 0x30, 0xb9, 0x91, 0xac, 0x5f, 0x22, 0xb9, 0x29, 0x3e, 0x61, 0xf3, 0x50, 0xfc,
	0x4b, 0x0b, 0x3e, 0x41, 0x32, 0x50, 0xdb, 0x93, 0x2d, 0x16, 0x5c, 0xbd, 0x22, 0x14, 0xb0, 0x52,
	0xba, 0x27, 0x57, 0xc4, 0x7c, 0x60, 0xeb, 0x0f, 0xbc, 0xe9, 0x9e, 0x4d, 0xc9, 0xb9, 0x25, 0x9e,
	0x61, 0x1d, 0x4a, 0x98, 0x43, 0xfc, 0x21, 0x9b, 0xeb, 0xea, 0x78, 0x92, 0xcd, 0xd1, 0x40, 0xfa,
	0x1e, 0xe2, 0xc1, 0x0a, 0x92, 0xd9, 0x32, 0xee, 0x58, 0xea, 0x2d, 0xa0, 0x3d, 0xde, 0xf1, 0x2d,
	0x8e, 0x7f, 0x58, 0x43, 0xe6, 0xed, 0x3e, 0xaa, 0x67, 0x09, 0x21, 0xa5, 0x7d, 0x1e, 0xd0, 0x7d,
	0xa8, 0x4b, 0xf8, 0x13, 0xea, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x70, 0x16, 0xc9, 0x97,
	0x01, 0x00, 0x00,
}
