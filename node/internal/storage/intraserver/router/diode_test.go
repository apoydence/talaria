//go:generate hel

package router_test

import (
	"sync"

	"github.com/apoydence/talaria/node/internal/storage/intraserver/router"
	"github.com/apoydence/talaria/pb/intra"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OneToOne", func() {
	var (
		d    *router.OneToOne
		data *intra.UpdateMessage

		mockAlerter *mockAlerter
	)

	Describe("Next()", func() {
		BeforeEach(func() {
			mockAlerter = newMockAlerter()

			d = router.NewOneToOne(5, mockAlerter)

			data = &intra.UpdateMessage{Name: "some-name"}
			d.Set(data)
		})

		It("returns the next data slice", func() {
			Expect(d.Next()).To(Equal(data))
		})

		Context("multiple data slices", func() {
			var (
				secondData *intra.UpdateMessage
			)

			BeforeEach(func() {
				secondData = &intra.UpdateMessage{Name: "some-other-name"}
				d.Set(secondData)
			})

			It("returns data slices in order", func() {
				Expect(d.Next()).To(Equal(data))
				Expect(d.Next()).To(Equal(secondData))
			})

			Describe("TryNext()", func() {
				It("returns true", func() {
					_, ok := d.TryNext()

					Expect(ok).To(BeTrue())
				})

				Context("reads exceed writes", func() {
					BeforeEach(func() {
						d.TryNext()
						d.TryNext()
					})

					It("returns false", func() {
						_, ok := d.TryNext()

						Expect(ok).To(BeFalse())
					})
				})
			})

			Context("reads exceed writes", func() {
				var (
					rxCh chan *intra.UpdateMessage
					wg   sync.WaitGroup
				)

				var waitForNext = func() {
					defer wg.Done()
					rxCh <- d.Next()
				}

				BeforeEach(func() {
					rxCh = make(chan *intra.UpdateMessage, 100)
					for i := 0; i < 2; i++ {
						d.Next()
					}
					wg.Add(1)
					go waitForNext()
				})

				AfterEach(func() {
					wg.Wait()
				})

				It("blocks until data is available", func() {
					Consistently(rxCh).Should(HaveLen(0))
					d.Set(data)
					Eventually(rxCh).Should(HaveLen(1))
				})
			})

			Context("buffer size exceeded", func() {
				BeforeEach(func() {
					for i := 0; i < 4; i++ {
						d.Set(secondData)
					}
				})

				It("wraps", func() {
					Expect(d.Next()).To(Equal(secondData))
				})

				It("alerts for each dropped point", func() {
					d.Next()
					Expect(mockAlerter.AlertInput.Missed).To(Receive(Equal(5)))
				})

				It("it updates the read index", func() {
					d.Next()
					Expect(mockAlerter.AlertInput.Missed).To(Receive(Equal(5)))

					for i := 0; i < 6; i++ {
						d.Set(secondData)
					}

					d.Next()
					Expect(mockAlerter.AlertInput.Missed).To(Receive(Equal(5)))
				})

				Context("read catches up with write", func() {
					BeforeEach(func() {
						d.Next()
						<-mockAlerter.AlertInput.Missed
					})

					It("does not alert", func() {
						d.Next()
						Expect(mockAlerter.AlertInput.Missed).To(Not(Receive()))
					})
				})

				Context("writer laps reader", func() {
					BeforeEach(func() {
						for i := 0; i < 5; i++ {
							d.Set(secondData)
						}
						d.Next()
					})

					It("sends an alert for each set", func() {
						Expect(mockAlerter.AlertInput.Missed).To(Receive(Equal(10)))
					})
				})
			})
		})
	})
})
