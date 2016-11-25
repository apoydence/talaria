package ringbuffer_test

import (
	"io"
	"testing"

	"github.com/apoydence/onpar"
	. "github.com/apoydence/onpar/expect"
	. "github.com/apoydence/onpar/matchers"
	"github.com/apoydence/talaria/node/internal/storage/buffers/ringbuffer"
	"github.com/coreos/etcd/raft/raftpb"
)

type TT struct {
	*testing.T
	d              *ringbuffer.RingBuffer
	valueA, valueB *raftpb.Entry
}

func TestRingBuffer(t *testing.T) {
	t.Parallel()

	o := onpar.New()
	defer o.Run(t)

	o.BeforeEach(func(t *testing.T) TT {
		d := ringbuffer.New(5)

		valueA := &raftpb.Entry{
			Data: []byte("some-value"),
		}
		valueB := &raftpb.Entry{
			Data: []byte("some-other-value"),
		}
		d.WriteTo(valueA)
		d.WriteTo(valueB)
		return TT{
			T:      t,
			d:      d,
			valueA: valueA,
			valueB: valueB,
		}
	})

	o.Group("WriteTo()", func() {
		o.Spec("it returns the written index", func(t TT) {
			idx, err := t.d.WriteTo(t.valueB)
			Expect(t, err == nil).To(Equal(true))
			Expect(t, idx).To(Equal(uint64(2)))
		})
	})

	o.Group("ReadAt()", func() {
		o.Spec("it returns expected value", func(t TT) {
			data, idx, err := t.d.ReadAt(0)
			Expect(t, err == nil).To(Equal(true))
			Expect(t, data).To(Equal(t.valueA))
			Expect(t, idx).To(Equal(uint64(0)))
		})

		o.Group("when reads exceed writes", func() {
			o.Spec("it returns io.EOF", func(t TT) {
				_, _, err := t.d.ReadAt(2)
				Expect(t, err).To(Equal(io.EOF))
			})
		})

		o.Group("LastIndex()", func() {
			o.Spec("it returns the last index", func(t TT) {
				Expect(t, t.d.LastIndex()).To(Equal(uint64(1)))
			})
		})

		o.Group("when buffer size exceeded", func() {
			o.BeforeEach(func(t TT) TT {
				for i := 0; i < 4; i++ {
					t.d.WriteTo(t.valueB)
				}

				return t
			})

			o.Spec("it wraps", func(t TT) {
				data, idx, err := t.d.ReadAt(0)
				Expect(t, err == nil).To(Equal(true))
				Expect(t, data).To(Equal(t.valueB))
				Expect(t, idx).To(Equal(uint64(5)))
			})
		})

	})

}
