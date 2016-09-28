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
		Arg0 chan context.Context
		Arg1 chan *intra.CreateInfo
	}
	CreateOutput struct {
		Ret0 chan *intra.CreateResponse
		Ret1 chan error
	}
}

func newMockIntraServer() *mockIntraServer {
	m := &mockIntraServer{}
	m.CreateCalled = make(chan bool, 100)
	m.CreateInput.Arg0 = make(chan context.Context, 100)
	m.CreateInput.Arg1 = make(chan *intra.CreateInfo, 100)
	m.CreateOutput.Ret0 = make(chan *intra.CreateResponse, 100)
	m.CreateOutput.Ret1 = make(chan error, 100)
	return m
}
func (m *mockIntraServer) Create(arg0 context.Context, arg1 *intra.CreateInfo) (*intra.CreateResponse, error) {
	m.CreateCalled <- true
	m.CreateInput.Arg0 <- arg0
	m.CreateInput.Arg1 <- arg1
	return <-m.CreateOutput.Ret0, <-m.CreateOutput.Ret1
}
