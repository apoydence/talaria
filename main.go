package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/apoydence/talaria/broker"
	"github.com/apoydence/talaria/kvstore"
	"github.com/apoydence/talaria/logging"
	"github.com/apoydence/talaria/orchestrator"
	"github.com/codegangsta/cli"
)

const (
	dataDir       = "dataDir"
	logLevel      = "logLevel"
	segmentLength = "segmentLength"
	numSegments   = "numSegments"
	port          = "port"
)

func main() {
	app := cli.NewApp()
	app.Name = "talaria"
	app.Usage = "Distribute your data"
	app.Action = func(c *cli.Context) {
		run(c)
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  dataDir + ", d",
			Usage: "The directory where the segmented files are stored",
		},
		cli.StringFlag{
			Name:  logLevel,
			Value: "INFO",
			Usage: "The log level",
		},
		cli.IntFlag{
			Name:  segmentLength + ", l",
			Value: 1024 * 1024,
			Usage: "The desired number of bytes for each segment",
		},
		cli.IntFlag{
			Name:  numSegments + ", n",
			Value: 1024 * 1024,
			Usage: "The desired number of bytes for each segment",
		},
		cli.IntFlag{
			Name:  port + ", p",
			Value: 8888,
			Usage: "The port to use",
		},
	}

	app.Run(os.Args)
}

func run(c *cli.Context) {
	validateFlags(c)
	setLogLevel(c)

	clientAddr := fmt.Sprintf("localhost:%d", c.Int(port))
	kvStore := kvstore.New(clientAddr)
	ioProvider := broker.NewFileProvider(c.String(dataDir), uint64(c.Int(segmentLength)), uint64(c.Int(numSegments)), time.Second)
	orch := orchestrator.New(clientAddr, ioWrapper(ioProvider.ProvideWriter), kvStore)

	broker.StartBrokerServer(c.Int(port), orch, ioProvider)
}

func setLogLevel(c *cli.Context) {
	logFlag := c.String(logLevel)
	var logLevel logging.LogLevel
	err := logLevel.UnmarshalJSON([]byte(logFlag))
	if err != nil {
		quit(fmt.Sprintf("Unable to parse log level: %s", logFlag), c)
	}

	logging.SetLevel(logLevel)
}

func validateFlags(c *cli.Context) {
	if !c.IsSet(dataDir) {
		quit(fmt.Sprintf("%s is required", dataDir), c)
	}
}

func quit(msg string, c *cli.Context) {
	fmt.Println(msg)
	cli.ShowAppHelp(c)
	os.Exit(1)
}

type ioWrapper func(name string) io.Writer

func (i ioWrapper) Add(name string) {
	i(name)
}