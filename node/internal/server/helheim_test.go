// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package server_test

import (
	"time"

	"github.com/poy/talaria/api/stored"
	"github.com/poy/talaria/node/internal/server"
)

type mockReader struct {
	ReadAtCalled chan bool
	ReadAtInput  struct {
		Index chan uint64
	}
	ReadAtOutput struct {
		Ret0 chan []byte
		Ret1 chan uint64
		Ret2 chan error
	}
	LastIndexCalled chan bool
	LastIndexOutput struct {
		Ret0 chan uint64
	}
}

func newMockReader() *mockReader {
	m := &mockReader{}
	m.ReadAtCalled = make(chan bool, 100)
	m.ReadAtInput.Index = make(chan uint64, 100)
	m.ReadAtOutput.Ret0 = make(chan []byte, 100)
	m.ReadAtOutput.Ret1 = make(chan uint64, 100)
	m.ReadAtOutput.Ret2 = make(chan error, 100)
	m.LastIndexCalled = make(chan bool, 100)
	m.LastIndexOutput.Ret0 = make(chan uint64, 100)
	return m
}
func (m *mockReader) ReadAt(index uint64) ([]byte, uint64, error) {
	m.ReadAtCalled <- true
	m.ReadAtInput.Index <- index
	return <-m.ReadAtOutput.Ret0, <-m.ReadAtOutput.Ret1, <-m.ReadAtOutput.Ret2
}
func (m *mockReader) LastIndex() uint64 {
	m.LastIndexCalled <- true
	return <-m.LastIndexOutput.Ret0
}

type mockWriter struct {
	WriteCalled chan bool
	WriteInput  struct {
		Data    chan stored.Data
		Timeout chan time.Duration
	}
	WriteOutput struct {
		Ret0 chan error
	}
}

func newMockWriter() *mockWriter {
	m := &mockWriter{}
	m.WriteCalled = make(chan bool, 100)
	m.WriteInput.Data = make(chan stored.Data, 100)
	m.WriteInput.Timeout = make(chan time.Duration, 100)
	m.WriteOutput.Ret0 = make(chan error, 100)
	return m
}
func (m *mockWriter) Write(data stored.Data, timeout time.Duration) error {
	m.WriteCalled <- true
	m.WriteInput.Data <- data
	m.WriteInput.Timeout <- timeout
	return <-m.WriteOutput.Ret0
}

type mockIOFetcher struct {
	FetchWriterCalled chan bool
	FetchWriterInput  struct {
		Name chan string
	}
	FetchWriterOutput struct {
		Ret0 chan server.Writer
		Ret1 chan error
	}
	FetchReaderCalled chan bool
	FetchReaderInput  struct {
		Name chan string
	}
	FetchReaderOutput struct {
		Ret0 chan server.Reader
		Ret1 chan error
	}
	FetchClustersCalled chan bool
	FetchClustersOutput struct {
		Ret0 chan []string
	}
}

func newMockIOFetcher() *mockIOFetcher {
	m := &mockIOFetcher{}
	m.FetchWriterCalled = make(chan bool, 100)
	m.FetchWriterInput.Name = make(chan string, 100)
	m.FetchWriterOutput.Ret0 = make(chan server.Writer, 100)
	m.FetchWriterOutput.Ret1 = make(chan error, 100)
	m.FetchReaderCalled = make(chan bool, 100)
	m.FetchReaderInput.Name = make(chan string, 100)
	m.FetchReaderOutput.Ret0 = make(chan server.Reader, 100)
	m.FetchReaderOutput.Ret1 = make(chan error, 100)
	m.FetchClustersCalled = make(chan bool, 100)
	m.FetchClustersOutput.Ret0 = make(chan []string, 100)
	return m
}
func (m *mockIOFetcher) FetchWriter(name string) (server.Writer, error) {
	m.FetchWriterCalled <- true
	m.FetchWriterInput.Name <- name
	return <-m.FetchWriterOutput.Ret0, <-m.FetchWriterOutput.Ret1
}
func (m *mockIOFetcher) FetchReader(name string) (server.Reader, error) {
	m.FetchReaderCalled <- true
	m.FetchReaderInput.Name <- name
	return <-m.FetchReaderOutput.Ret0, <-m.FetchReaderOutput.Ret1
}
func (m *mockIOFetcher) FetchClusters() []string {
	m.FetchClustersCalled <- true
	return <-m.FetchClustersOutput.Ret0
}
