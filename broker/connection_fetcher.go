package broker

import (
	"fmt"
	"net/url"
	"sync"

	"github.com/apoydence/talaria/logging"
)

var (
	BlacklistedErr = fmt.Errorf("blacklisted")
)

type ConnectionFetcher struct {
	log       logging.Logger
	blacklist []string

	syncConns    sync.RWMutex
	conns        []*Connection
	nextFetchIdx uint64
}

func NewConnectionFetcher(blacklist []string, URLs ...string) (*ConnectionFetcher, error) {
	log := logging.Log("ConnectionFetcher")
	log.Debug("Broker List: %v", URLs)

	var conns []*Connection
	for _, URL := range URLs {
		verifyUrl(URL, log)
		conn, err := NewConnection(URL)
		if err != nil {
			return nil, err
		}
		conns = append(conns, conn)
	}

	return &ConnectionFetcher{
		log:       log,
		blacklist: blacklist,
		conns:     conns,
	}, nil
}

func (c *ConnectionFetcher) Fetch(fileName string, create bool) (*Connection, uint64, error) {
	c.syncConns.Lock()
	defer c.syncConns.Unlock()

	fileId := c.getNextFetchIdx()
	conn := c.verifyConn(c.roundRobinConns(fileId))

	err := conn.FetchFile(fileId, fileName, create)
	if err == nil {
		return c.checkBlacklist(conn, fileId, nil)
	}

	if err.Uri == "" {
		return nil, 0, fmt.Errorf(err.errMessage)
	}

	conn = c.verifyConn(c.fetchConnection(err.Uri))
	if conn == nil {
		conn = c.createConnection(err.Uri)
		c.conns = append(c.conns, conn)
	}

	err = conn.FetchFile(fileId, fileName, create)
	return c.checkBlacklist(conn, fileId, err)
}

func (c *ConnectionFetcher) Close() {
	for _, conn := range c.conns {
		conn.Close()
	}
}

func (c *ConnectionFetcher) getNextFetchIdx() uint64 {
	defer func() {
		c.nextFetchIdx++
	}()
	return c.nextFetchIdx
}

func (c *ConnectionFetcher) createConnection(URL string) *Connection {
	conn, err := NewConnection(URL)
	if err != nil {
		c.log.Panicf("Unable to create connection to %s: %v", URL, err)
	}

	return conn
}

func (c *ConnectionFetcher) verifyConn(conn *Connection) *Connection {
	if !conn.Errored() {
		return conn
	}

	index, _ := c.fetchConnectionIndex(conn.URL)
	c.conns = append(c.conns[:index], c.conns[index+1:]...)

	newConn := c.createConnection(conn.URL)
	c.conns = append(c.conns, newConn)

	return newConn
}

func (c *ConnectionFetcher) roundRobinConns(fileId uint64) *Connection {
	return c.conns[int(fileId)%len(c.conns)]
}

func (c *ConnectionFetcher) fetchConnection(URL string) *Connection {
	_, conn := c.fetchConnectionIndex(URL)
	return conn
}

func (c *ConnectionFetcher) fetchConnectionIndex(URL string) (int, *Connection) {
	c.log.Debug("Fetching connection for %s", URL)
	for i, conn := range c.conns {
		if conn.URL == URL {
			c.log.Debug("Found connection for %s", URL)
			return i, conn
		}
	}
	return 0, nil
}

func (c *ConnectionFetcher) checkBlacklist(conn *Connection, fileId uint64, err *ConnectionError) (*Connection, uint64, error) {
	if err != nil {
		return nil, 0, fmt.Errorf(err.Error())
	}

	if c.onBlacklist(conn) {
		return nil, 0, BlacklistedErr
	}

	return conn, fileId, nil
}

func (c *ConnectionFetcher) onBlacklist(conn *Connection) bool {
	for _, URL := range c.blacklist {
		if URL == conn.URL {
			return true
		}
	}

	return false
}

func verifyUrl(URL string, log logging.Logger) {
	u, _ := url.Parse(URL)
	if u == nil || u.Host == "" {
		log.Panicf("Invalid URL: %s", URL)
	}
}
