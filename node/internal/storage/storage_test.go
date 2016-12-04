//go:generate hel

package storage_test

import (
	"testing"

	"github.com/apoydence/onpar"
	. "github.com/apoydence/onpar/expect"
	. "github.com/apoydence/onpar/matchers"
	"github.com/apoydence/talaria/node/internal/server"
	"github.com/apoydence/talaria/node/internal/storage"
	"github.com/apoydence/talaria/pb/intra"
	"github.com/coreos/etcd/raft/raftpb"
)

type TT struct {
	*testing.T
	mockURIFinder *mockURIFinder
	mockReceiver  *mockReceiver
	callback      func() (raftpb.Message, error)
	callbackMsgs  chan raftpb.Message
	callbackErrs  chan error
	fetcher       *storage.Storage
	readerA       server.Reader
	peers         []*intra.PeerInfo
}

func TestStorageFetchReader(t *testing.T) {
	t.Parallel()
	o := onpar.New()
	defer o.Run(t)

	o.BeforeEach(func(t *testing.T) TT {
		peers := []*intra.PeerInfo{
			{
				Uri: "some-uri",
				Id:  99,
			},
		}

		mockURIFinder := newMockURIFinder()

		mockReceiver := newMockReceiver()
		callbackMsgs := make(chan raftpb.Message, 100)
		callbackErrs := make(chan error, 100)
		callback := func() (raftpb.Message, error) {
			return <-callbackMsgs, <-callbackErrs
		}
		mockReceiver.ReceiverOutput.Ret0 <- callback

		return TT{
			T:             t,
			mockReceiver:  mockReceiver,
			mockURIFinder: mockURIFinder,
			fetcher:       storage.New(99, mockReceiver, mockURIFinder),
			peers:         peers,
			callback:      callback,
			callbackMsgs:  callbackMsgs,
			callbackErrs:  callbackErrs,
		}
	})

	o.Group("when Create() has been called", func() {
		o.BeforeEach(func(t TT) TT {
			err := t.fetcher.Create("some-buffer", t.peers)
			Expect(t, err == nil).To(Equal(true))
			return t
		})

		o.Spec("it returns the same reader each time", func(t TT) {
			readerA, err := t.fetcher.FetchReader("some-buffer")
			Expect(t, err == nil).To(Equal(true))

			readerB, err := t.fetcher.FetchReader("some-buffer")
			Expect(t, err == nil).To(Equal(true))

			Expect(t, readerA != nil).To(Equal(true))
			Expect(t, readerA).To(Equal(readerB))
		})

		o.Group("when Create() has been called twice for the same buffer", func() {
			o.BeforeEach(func(t TT) TT {
				readerA, _ := t.fetcher.FetchReader("some-buffer")
				err := t.fetcher.Create("some-buffer", t.peers)
				Expect(t, err == nil).To(Equal(true))
				t.readerA = readerA
				return t
			})

			o.Spec("it still returns the same reader each time", func(t TT) {
				readerB, _ := t.fetcher.FetchReader("some-buffer")

				Expect(t, t.readerA).To(Equal(readerB))
			})
		})
	})

	o.Group("when Create() has not been called", func() {
		o.Spec("it returns an error", func(t TT) {
			_, err := t.fetcher.FetchReader("some-buffer")
			Expect(t, err != nil).To(Equal(true))
		})
	})
}

func TestStorageFetchWriter(t *testing.T) {
	t.Parallel()
	o := onpar.New()
	defer o.Run(t)

	o.BeforeEach(func(t *testing.T) TT {
		peers := []*intra.PeerInfo{
			{
				Uri: "some-uri",
				Id:  99,
			},
		}

		mockURIFinder := newMockURIFinder()

		mockReceiver := newMockReceiver()
		callbackMsgs := make(chan raftpb.Message, 100)
		callbackErrs := make(chan error, 100)
		callback := func() (raftpb.Message, error) {
			return <-callbackMsgs, <-callbackErrs
		}
		mockReceiver.ReceiverOutput.Ret0 <- callback

		return TT{
			T:             t,
			mockURIFinder: mockURIFinder,
			mockReceiver:  mockReceiver,
			fetcher:       storage.New(99, mockReceiver, mockURIFinder),
			peers:         peers,
			callback:      callback,
		}
	})

	o.Group("when Create() has been called", func() {
		o.BeforeEach(func(t TT) TT {
			err := t.fetcher.Create("some-buffer", t.peers)
			Expect(t, err == nil).To(Equal(true))

			return t
		})

		o.Spec("it returns the same writer each time", func(t TT) {
			writerA, err := t.fetcher.FetchWriter("some-buffer")
			Expect(t, err == nil).To(Equal(true))

			writerB, err := t.fetcher.FetchWriter("some-buffer")
			Expect(t, err == nil).To(Equal(true))

			Expect(t, writerA != nil).To(Equal(true))
			Expect(t, writerA).To(Equal(writerB))
		})
	})

	o.Group("when Create() has not been called", func() {
		o.Spec("it returns an error", func(t TT) {
			_, err := t.fetcher.FetchWriter("some-buffer")
			Expect(t, err != nil).To(Equal(true))
		})
	})
}

func TestStorageLeader(t *testing.T) {
	t.Parallel()
	o := onpar.New()
	defer o.Run(t)

	o.BeforeEach(func(t *testing.T) TT {
		peers := []*intra.PeerInfo{
			{
				Uri: "some-uri",
				Id:  99,
			},
		}

		mockURIFinder := newMockURIFinder()

		mockReceiver := newMockReceiver()
		callbackMsgs := make(chan raftpb.Message, 100)
		callbackErrs := make(chan error, 100)
		callback := func() (raftpb.Message, error) {
			return <-callbackMsgs, <-callbackErrs
		}
		mockReceiver.ReceiverOutput.Ret0 <- callback

		return TT{
			T:             t,
			mockURIFinder: mockURIFinder,
			mockReceiver:  mockReceiver,
			fetcher:       storage.New(99, mockReceiver, mockURIFinder),
			peers:         peers,
			callback:      callback,
		}
	})

	o.Group("when Create() has been called", func() {
		o.BeforeEach(func(t TT) TT {
			err := t.fetcher.Create("some-buffer", t.peers)
			Expect(t, err == nil).To(Equal(true))
			return t
		})

		o.Spec("", func(t TT) {
		})
	})

}