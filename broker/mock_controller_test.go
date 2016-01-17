package broker_test

import "github.com/apoydence/talaria/broker"

type mockController struct {
	fetchFileCh       chan string
	fetchFileIdCh     chan uint64
	fetchFileErrCh    chan *broker.ConnectionError
	fetchFileCreateCh chan bool

	writeFileIdCh    chan uint64
	writeFileDataCh  chan []byte
	writeFileIndexCh chan int64
	writeFileErrCh   chan error

	readFileIdCh   chan uint64
	readFileDataCh chan []byte
	readIndexCh    chan int64
	readFileErrCh  chan error

	initIdCh       chan uint64
	initDataCh     chan []byte
	initIndexCh    chan int64
	initIndexRetCh chan int64
	initErrCh      chan error

	seekIdCh       chan uint64
	seekIndexCh    chan uint64
	seekCallbackCh chan func(error)

	validLeaderNameCh   chan string
	validLeaderIndexCh  chan uint64
	validLeaderResultCh chan bool
}

func newMockController() *mockController {
	return &mockController{
		fetchFileCh:       make(chan string, 100),
		fetchFileIdCh:     make(chan uint64, 100),
		fetchFileErrCh:    make(chan *broker.ConnectionError, 100),
		fetchFileCreateCh: make(chan bool, 100),

		writeFileIdCh:    make(chan uint64, 100),
		writeFileDataCh:  make(chan []byte, 100),
		writeFileIndexCh: make(chan int64, 100),
		writeFileErrCh:   make(chan error, 100),

		readFileIdCh:   make(chan uint64, 100),
		readFileDataCh: make(chan []byte, 100),
		readIndexCh:    make(chan int64, 100),
		readFileErrCh:  make(chan error, 100),

		initIdCh:       make(chan uint64, 100),
		initDataCh:     make(chan []byte, 100),
		initIndexCh:    make(chan int64, 100),
		initIndexRetCh: make(chan int64, 100),
		initErrCh:      make(chan error, 100),

		seekIdCh:       make(chan uint64, 100),
		seekIndexCh:    make(chan uint64, 100),
		seekCallbackCh: make(chan func(error), 100),

		validLeaderNameCh:   make(chan string, 100),
		validLeaderIndexCh:  make(chan uint64, 100),
		validLeaderResultCh: make(chan bool, 100),
	}
}

func (m *mockController) FetchFile(fileId uint64, name string, create bool) *broker.ConnectionError {
	m.fetchFileCh <- name
	m.fetchFileIdCh <- fileId
	m.fetchFileCreateCh <- create
	return <-m.fetchFileErrCh
}

func (m *mockController) WriteToFile(id uint64, data []byte) (int64, error) {
	m.writeFileIdCh <- id
	m.writeFileDataCh <- data
	return <-m.writeFileIndexCh, <-m.writeFileErrCh
}

func (m *mockController) ReadFromFile(id uint64, callback func([]byte, int64, error)) {
	m.readFileIdCh <- id
	go callback(<-m.readFileDataCh, <-m.readIndexCh, <-m.readFileErrCh)
}

func (m *mockController) SeekIndex(id, index uint64, callback func(error)) {
	m.seekIdCh <- id
	m.seekIndexCh <- index
	m.seekCallbackCh <- callback
}

func (m *mockController) InitWriteIndex(id uint64, index int64, data []byte) (int64, error) {
	m.initIdCh <- id
	m.initDataCh <- data
	m.initIndexCh <- index
	return <-m.initIndexRetCh, <-m.initErrCh
}

func (m *mockController) ValidateLeader(name string, index uint64) bool {
	m.validLeaderNameCh <- name
	m.validLeaderIndexCh <- index
	return <-m.validLeaderResultCh
}

func (m *mockController) closeChannels() {
	close(m.fetchFileCh)
	close(m.fetchFileIdCh)
	close(m.fetchFileErrCh)

	close(m.writeFileIdCh)
	close(m.writeFileDataCh)
	close(m.writeFileIndexCh)
	close(m.writeFileErrCh)

	close(m.readFileIdCh)
	close(m.readFileDataCh)
	close(m.readFileErrCh)
}
