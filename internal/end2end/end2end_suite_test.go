package end2end_test

import (
	"fmt"
	"os/exec"
	"time"

	"google.golang.org/grpc"

	"github.com/apoydence/talaria/internal/end2end"
	"github.com/apoydence/talaria/pb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

func TestEnd2end(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "End2end Suite")
}

var (
	nodePort      int
	schedulerPort int

	sessions    []*gexec.Session
	clientConns []*grpc.ClientConn

	nodeClient      pb.TalariaClient
	schedulerClient pb.SchedulerClient
)

var _ = BeforeSuite(func() {
	startNode()
	startScheduler()
	nodeClient = connectToNode()
	schedulerClient = connectToScheduler()
})

var _ = AfterSuite(func() {
	for _, clientConn := range clientConns {
		clientConn.Close()
	}

	for _, session := range sessions {
		session.Kill().Wait(5 * time.Second)
	}

	gexec.CleanupBuildArtifacts()
})

func connectToNode() pb.TalariaClient {
	clientConn, err := grpc.Dial(fmt.Sprintf("localhost:%d", nodePort), grpc.WithInsecure())
	Expect(err).ToNot(HaveOccurred())
	clientConns = append(clientConns, clientConn)

	return pb.NewTalariaClient(clientConn)
}

func connectToScheduler() pb.SchedulerClient {
	clientConn, err := grpc.Dial(fmt.Sprintf("localhost:%d", schedulerPort), grpc.WithInsecure())
	Expect(err).ToNot(HaveOccurred())
	clientConns = append(clientConns, clientConn)

	return pb.NewSchedulerClient(clientConn)
}

func startNode() {
	nodePort = end2end.AvailablePort()
	path, err := gexec.Build("github.com/apoydence/talaria/node")
	Expect(err).ToNot(HaveOccurred())
	command := exec.Command(path)
	command.Env = []string{
		fmt.Sprintf("PORT=%d", nodePort),
	}

	nodeSession, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	Consistently(nodeSession.Exited).ShouldNot(BeClosed())
	sessions = append(sessions, nodeSession)
}

func startScheduler() {
	schedulerPort = end2end.AvailablePort()
	path, err := gexec.Build("github.com/apoydence/talaria/scheduler")
	Expect(err).ToNot(HaveOccurred())
	command := exec.Command(path)
	command.Env = []string{
		fmt.Sprintf("PORT=%d", schedulerPort),
		fmt.Sprintf("NODES=localhost:%d", nodePort),
	}

	schedulerSession, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	Consistently(schedulerSession.Exited).ShouldNot(BeClosed())
	sessions = append(sessions, schedulerSession)
}
