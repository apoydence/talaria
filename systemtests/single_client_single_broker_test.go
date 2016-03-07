package systemtests_test

import (
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/apoydence/talaria/client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("SingleConnectionSingleBroker", func() {

	var (
		session *gexec.Session
		client  *client.Client
		URL     string
	)

	BeforeEach(func() {
		var err error
		tmpDir, err = ioutil.TempDir("/tmp", "systemtalaria")
		Expect(err).ToNot(HaveOccurred())
		URL, session = startTalaria(tmpDir)
		client = startClient(URL)
	})

	AfterEach(func() {
		session.Kill()
		session.Wait("10s", "100ms")

		Expect(os.RemoveAll(tmpDir)).To(Succeed())
		client.Close()
	})

	It("writes and reads from a single file", func(done Done) {
		defer close(done)
		name := "some-file"

		Expect(client.CreateFile(name)).To(Succeed())
		writer, err := client.FetchWriter(name)
		Expect(err).ToNot(HaveOccurred())

		for i := byte(0); i < 100; i++ {
			_, err := writer.WriteToFile([]byte{i})
			Expect(err).ToNot(HaveOccurred())
		}

		reader, err := client.FetchReader(name)
		Expect(err).ToNot(HaveOccurred())

		for i := 0; i < 100; i++ {
			data, index, err := reader.ReadFromFile()
			Expect(err).ToNot(HaveOccurred())
			Expect(data).To(HaveLen(1))
			Expect(data[0]).To(BeEquivalentTo(i))
			Expect(index).To(BeEquivalentTo(i))
		}
	}, 5)

	It("seeks in a single file", func(done Done) {
		defer close(done)

		name := "some-file"

		Expect(client.CreateFile(name)).To(Succeed())
		writer, err := client.FetchWriter(name)
		Expect(err).ToNot(HaveOccurred())

		for i := byte(0); i < 100; i++ {
			_, err := writer.WriteToFile([]byte{i})
			Expect(err).ToNot(HaveOccurred())
		}

		reader, err := client.FetchReader(name)
		Expect(err).ToNot(HaveOccurred())
		By("Seeking to 50")
		Expect(reader.SeekIndex(50)).To(Succeed())
		By("Done seeking to 50")

		for i := 50; i < 100; i++ {
			data, index, err := reader.ReadFromFile()
			Expect(err).ToNot(HaveOccurred())
			Expect(data).To(HaveLen(1))
			Expect(data[0]).To(BeEquivalentTo(i))
			Expect(index).To(BeEquivalentTo(i))
		}
	}, 5)

	It("writes and reads from a single file at the same time", func(done Done) {
		defer close(done)
		name := "some-file"
		clientW := startClient(URL)
		clientR := startClient(URL)

		Expect(client.CreateFile(name)).To(Succeed())

		var wg sync.WaitGroup
		defer wg.Wait()
		wg.Add(1)
		go func() {
			defer GinkgoRecover()
			defer wg.Done()

			writer, err := clientW.FetchWriter(name)
			Expect(err).ToNot(HaveOccurred())

			for i := 0; i < 10; i++ {
				_, err := writer.WriteToFile([]byte{byte(i)})
				Expect(err).ToNot(HaveOccurred())
				time.Sleep(time.Millisecond)
			}
		}()

		reader, err := clientR.FetchReader(name)
		Expect(err).ToNot(HaveOccurred())

		var result []byte
		for len(result) < 10 {
			data, _, err := reader.ReadFromFile()
			Expect(err).ToNot(HaveOccurred())
			result = append(result, data...)
		}

		for i := 0; i < 10; i++ {
			Expect(result[i]).To(BeEquivalentTo(i))
		}
	}, 5)

	It("reports the leader via the FileMeta()", func(done Done) {
		defer close(done)

		name := "some-file"

		client := startClient(URL)

		Expect(client.CreateFile(name)).To(Succeed())
		meta, err := client.FileMeta(name)
		Expect(err).ToNot(HaveOccurred())
		Expect(meta.GetReplicaURIs()).To(HaveLen(1))
		Expect(meta.GetReplicaURIs()[0]).To(Equal(URL))
	}, 5)

})
