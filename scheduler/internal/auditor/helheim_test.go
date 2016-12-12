// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package auditor_test

import (
	"time"

	"github.com/apoydence/talaria/pb/intra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type mockNode struct {
	StatusCalled chan bool
	StatusInput  struct {
		Ctx  chan context.Context
		Req  chan *intra.StatusRequest
		Opts chan []grpc.CallOption
	}
	StatusOutput struct {
		Ret0 chan *intra.StatusResponse
		Ret1 chan error
	}
	CreateCalled chan bool
	CreateInput  struct {
		Ctx  chan context.Context
		Req  chan *intra.CreateInfo
		Opts chan []grpc.CallOption
	}
	CreateOutput struct {
		Ret0 chan *intra.CreateResponse
		Ret1 chan error
	}
	UpdateConfigCalled chan bool
	UpdateConfigInput  struct {
		Ctx  chan context.Context
		Req  chan *intra.UpdateConfigRequest
		Opts chan []grpc.CallOption
	}
	UpdateConfigOutput struct {
		Ret0 chan *intra.UpdateConfigResponse
		Ret1 chan error
	}
	LeaderCalled chan bool
	LeaderInput  struct {
		Ctx  chan context.Context
		Req  chan *intra.LeaderRequest
		Opts chan []grpc.CallOption
	}
	LeaderOutput struct {
		Ret0 chan *intra.LeaderInfo
		Ret1 chan error
	}
}

func newMockNode() *mockNode {
	m := &mockNode{}
	m.StatusCalled = make(chan bool, 100)
	m.StatusInput.Ctx = make(chan context.Context, 100)
	m.StatusInput.Req = make(chan *intra.StatusRequest, 100)
	m.StatusInput.Opts = make(chan []grpc.CallOption, 100)
	m.StatusOutput.Ret0 = make(chan *intra.StatusResponse, 100)
	m.StatusOutput.Ret1 = make(chan error, 100)
	m.CreateCalled = make(chan bool, 100)
	m.CreateInput.Ctx = make(chan context.Context, 100)
	m.CreateInput.Req = make(chan *intra.CreateInfo, 100)
	m.CreateInput.Opts = make(chan []grpc.CallOption, 100)
	m.CreateOutput.Ret0 = make(chan *intra.CreateResponse, 100)
	m.CreateOutput.Ret1 = make(chan error, 100)
	m.UpdateConfigCalled = make(chan bool, 100)
	m.UpdateConfigInput.Ctx = make(chan context.Context, 100)
	m.UpdateConfigInput.Req = make(chan *intra.UpdateConfigRequest, 100)
	m.UpdateConfigInput.Opts = make(chan []grpc.CallOption, 100)
	m.UpdateConfigOutput.Ret0 = make(chan *intra.UpdateConfigResponse, 100)
	m.UpdateConfigOutput.Ret1 = make(chan error, 100)
	m.LeaderCalled = make(chan bool, 100)
	m.LeaderInput.Ctx = make(chan context.Context, 100)
	m.LeaderInput.Req = make(chan *intra.LeaderRequest, 100)
	m.LeaderInput.Opts = make(chan []grpc.CallOption, 100)
	m.LeaderOutput.Ret0 = make(chan *intra.LeaderInfo, 100)
	m.LeaderOutput.Ret1 = make(chan error, 100)
	return m
}
func (m *mockNode) Status(ctx context.Context, req *intra.StatusRequest, opts ...grpc.CallOption) (*intra.StatusResponse, error) {
	m.StatusCalled <- true
	m.StatusInput.Ctx <- ctx
	m.StatusInput.Req <- req
	m.StatusInput.Opts <- opts
	return <-m.StatusOutput.Ret0, <-m.StatusOutput.Ret1
}
func (m *mockNode) Create(ctx context.Context, req *intra.CreateInfo, opts ...grpc.CallOption) (*intra.CreateResponse, error) {
	m.CreateCalled <- true
	m.CreateInput.Ctx <- ctx
	m.CreateInput.Req <- req
	m.CreateInput.Opts <- opts
	return <-m.CreateOutput.Ret0, <-m.CreateOutput.Ret1
}
func (m *mockNode) UpdateConfig(ctx context.Context, req *intra.UpdateConfigRequest, opts ...grpc.CallOption) (*intra.UpdateConfigResponse, error) {
	m.UpdateConfigCalled <- true
	m.UpdateConfigInput.Ctx <- ctx
	m.UpdateConfigInput.Req <- req
	m.UpdateConfigInput.Opts <- opts
	return <-m.UpdateConfigOutput.Ret0, <-m.UpdateConfigOutput.Ret1
}
func (m *mockNode) Leader(ctx context.Context, req *intra.LeaderRequest, opts ...grpc.CallOption) (*intra.LeaderInfo, error) {
	m.LeaderCalled <- true
	m.LeaderInput.Ctx <- ctx
	m.LeaderInput.Req <- req
	m.LeaderInput.Opts <- opts
	return <-m.LeaderOutput.Ret0, <-m.LeaderOutput.Ret1
}

type mockContext struct {
	DeadlineCalled chan bool
	DeadlineOutput struct {
		Deadline chan time.Time
		Ok       chan bool
	}
	DoneCalled chan bool
	DoneOutput struct {
		Ret0 chan (<-chan struct{})
	}
	ErrCalled chan bool
	ErrOutput struct {
		Ret0 chan error
	}
	ValueCalled chan bool
	ValueInput  struct {
		Key chan interface{}
	}
	ValueOutput struct {
		Ret0 chan interface{}
	}
}

func newMockContext() *mockContext {
	m := &mockContext{}
	m.DeadlineCalled = make(chan bool, 100)
	m.DeadlineOutput.Deadline = make(chan time.Time, 100)
	m.DeadlineOutput.Ok = make(chan bool, 100)
	m.DoneCalled = make(chan bool, 100)
	m.DoneOutput.Ret0 = make(chan (<-chan struct{}), 100)
	m.ErrCalled = make(chan bool, 100)
	m.ErrOutput.Ret0 = make(chan error, 100)
	m.ValueCalled = make(chan bool, 100)
	m.ValueInput.Key = make(chan interface{}, 100)
	m.ValueOutput.Ret0 = make(chan interface{}, 100)
	return m
}
func (m *mockContext) Deadline() (deadline time.Time, ok bool) {
	m.DeadlineCalled <- true
	return <-m.DeadlineOutput.Deadline, <-m.DeadlineOutput.Ok
}
func (m *mockContext) Done() <-chan struct{} {
	m.DoneCalled <- true
	return <-m.DoneOutput.Ret0
}
func (m *mockContext) Err() error {
	m.ErrCalled <- true
	return <-m.ErrOutput.Ret0
}
func (m *mockContext) Value(key interface{}) interface{} {
	m.ValueCalled <- true
	m.ValueInput.Key <- key
	return <-m.ValueOutput.Ret0
}