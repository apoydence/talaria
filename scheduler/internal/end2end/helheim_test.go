// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package end2end_test

import (
	"golang.org/x/net/context"

	"github.com/apoydence/talaria/pb/intra"
)

type mockIntraServer struct {
	CreateCalled chan bool
	CreateInput  struct {
		Ctx chan context.Context
		In  chan *intra.CreateInfo
	}
	CreateOutput struct {
		Ret0 chan *intra.CreateResponse
		Ret1 chan error
	}
	LeaderCalled chan bool
	LeaderInput  struct {
		Ctx chan context.Context
		In  chan *intra.LeaderRequest
	}
	LeaderOutput struct {
		Ret0 chan *intra.LeaderInfo
		Ret1 chan error
	}
	StatusCalled chan bool
	StatusInput  struct {
		Ctx chan context.Context
		In  chan *intra.StatusRequest
	}
	StatusOutput struct {
		Ret0 chan *intra.StatusResponse
		Ret1 chan error
	}
	UpdateCalled chan bool
	UpdateInput  struct {
		Ctx chan context.Context
		In  chan *intra.UpdateMessage
	}
	UpdateOutput struct {
		Ret0 chan *intra.UpdateResponse
		Ret1 chan error
	}
}

func newMockIntraServer() *mockIntraServer {
	m := &mockIntraServer{}
	m.CreateCalled = make(chan bool, 100)
	m.CreateInput.Ctx = make(chan context.Context, 100)
	m.CreateInput.In = make(chan *intra.CreateInfo, 100)
	m.CreateOutput.Ret0 = make(chan *intra.CreateResponse, 100)
	m.CreateOutput.Ret1 = make(chan error, 100)
	m.LeaderCalled = make(chan bool, 100)
	m.LeaderInput.Ctx = make(chan context.Context, 100)
	m.LeaderInput.In = make(chan *intra.LeaderRequest, 100)
	m.LeaderOutput.Ret0 = make(chan *intra.LeaderInfo, 100)
	m.LeaderOutput.Ret1 = make(chan error, 100)
	m.StatusCalled = make(chan bool, 100)
	m.StatusInput.Ctx = make(chan context.Context, 100)
	m.StatusInput.In = make(chan *intra.StatusRequest, 100)
	m.StatusOutput.Ret0 = make(chan *intra.StatusResponse, 100)
	m.StatusOutput.Ret1 = make(chan error, 100)
	m.UpdateCalled = make(chan bool, 100)
	m.UpdateInput.Ctx = make(chan context.Context, 100)
	m.UpdateInput.In = make(chan *intra.UpdateMessage, 100)
	m.UpdateOutput.Ret0 = make(chan *intra.UpdateResponse, 100)
	m.UpdateOutput.Ret1 = make(chan error, 100)
	return m
}
func (m *mockIntraServer) Create(ctx context.Context, in *intra.CreateInfo) (*intra.CreateResponse, error) {
	m.CreateCalled <- true
	m.CreateInput.Ctx <- ctx
	m.CreateInput.In <- in
	return <-m.CreateOutput.Ret0, <-m.CreateOutput.Ret1
}
func (m *mockIntraServer) Leader(ctx context.Context, in *intra.LeaderRequest) (*intra.LeaderInfo, error) {
	m.LeaderCalled <- true
	m.LeaderInput.Ctx <- ctx
	m.LeaderInput.In <- in
	return <-m.LeaderOutput.Ret0, <-m.LeaderOutput.Ret1
}
func (m *mockIntraServer) Status(ctx context.Context, in *intra.StatusRequest) (*intra.StatusResponse, error) {
	m.StatusCalled <- true
	m.StatusInput.Ctx <- ctx
	m.StatusInput.In <- in
	return <-m.StatusOutput.Ret0, <-m.StatusOutput.Ret1
}
func (m *mockIntraServer) Update(ctx context.Context, in *intra.UpdateMessage) (*intra.UpdateResponse, error) {
	m.UpdateCalled <- true
	m.UpdateInput.Ctx <- ctx
	m.UpdateInput.In <- in
	return <-m.UpdateOutput.Ret0, <-m.UpdateOutput.Ret1
}
