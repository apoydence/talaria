package systemtests_test

import (
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/apoydence/talaria/broker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("SingleClientSingleBroker", func() {

	var (
		URL     string
		session *gexec.Session
	)

	BeforeEach(func() {
		var err error
		tmpDir, err = ioutil.TempDir("/tmp", "systemtalaria")
		Expect(err).ToNot(HaveOccurred())
		URL, session = startTalaria(tmpDir)
	})

	AfterEach(func() {
		session.Kill()
		session.Wait("10s", "100ms")

		Expect(os.RemoveAll(tmpDir)).To(Succeed())
	})

	It("Writes and reads from a single file", func() {
		client := broker.NewClient(URL)
		fileId, err := client.FetchFile("some-file")
		Expect(err).ToNot(HaveOccurred())
		for i := byte(0); i < 100; i++ {
			_, err = client.WriteToFile(fileId, []byte{i})
			Expect(err).ToNot(HaveOccurred())
		}

		data, err := client.ReadFromFile(fileId)
		Expect(err).ToNot(HaveOccurred())
		for i := 0; i < 100; i++ {
			Expect(data[i]).To(BeEquivalentTo(i))
		}
	})

	It("Writes and reads from a single file at the same time", func() {
		clientW := broker.NewClient(URL)
		clientR := broker.NewClient(URL)
		fileIdW, err := clientW.FetchFile("some-file")
		Expect(err).ToNot(HaveOccurred())
		fileIdR, err := clientR.FetchFile("some-file")
		Expect(err).ToNot(HaveOccurred())

		wg := sync.WaitGroup{}
		defer wg.Wait()
		wg.Add(1)
		go func() {
			defer GinkgoRecover()
			defer wg.Done()
			for i := 0; i < 10; i++ {
				_, err = clientW.WriteToFile(fileIdW, []byte{byte(i)})
				time.Sleep(time.Millisecond)
				Expect(err).ToNot(HaveOccurred())
			}
		}()

		var result []byte
		for len(result) < 10 {
			data, err := clientR.ReadFromFile(fileIdR)
			Expect(err).ToNot(HaveOccurred())
			result = append(result, data...)
		}

		for i := 0; i < 10; i++ {
			Expect(result[i]).To(BeEquivalentTo(i))
		}
	})

})