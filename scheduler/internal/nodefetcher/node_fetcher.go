package nodefetcher

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/apoydence/talaria/pb/intra"
	"github.com/apoydence/talaria/scheduler/internal/server"
	"google.golang.org/grpc"
)

type NodeFetcher struct {
	clients []clientInfo
}

type clientInfo struct {
	URI    string
	conn   *grpc.ClientConn
	client intra.NodeClient
}

func New(URIs []string) *NodeFetcher {
	var clients []clientInfo
	for _, URI := range URIs {
		conn, err := grpc.Dial(URI, grpc.WithInsecure())
		if err != nil {
			log.Panicf("Error connecting to %s: %s", URI, err)
		}
		client := intra.NewNodeClient(conn)
		clients = append(clients, clientInfo{
			conn:   conn,
			client: client,
			URI:    URI,
		})
	}

	return &NodeFetcher{
		clients: clients,
	}
}

func (f *NodeFetcher) FetchAllNodes() []server.NodeInfo {
	var results []server.NodeInfo
	for _, client := range f.clients {
		id, ok := f.fetchID(client, 0)
		if !ok {
			continue
		}

		results = append(results, server.NodeInfo{
			Client: client.client,
			URI:    client.URI,
			ID:     id,
		})
	}
	return results
}

func (f *NodeFetcher) FetchNodes(count int, exclude ...server.NodeInfo) []server.NodeInfo {
	if len(f.clients) < count {
		return nil
	}

	clients := f.exclude(exclude)

	num := f.permCount(len(clients), count)

	var infos []server.NodeInfo
	for _, p := range rand.Perm(num) {
		id, ok := f.fetchID(clients[p], 0)
		if !ok {
			continue
		}

		infos = append(infos, server.NodeInfo{
			Client: clients[p].client,
			URI:    clients[p].URI,
			ID:     id,
		})
	}

	return infos
}

func (f *NodeFetcher) permCount(max, count int) int {
	if max < count {
		return max
	}

	return count
}

func (f *NodeFetcher) exclude(exclude []server.NodeInfo) []clientInfo {
	var result []clientInfo
	for _, n := range f.clients {
		if !f.excluded(n, exclude) {
			result = append(result, n)
		}
	}

	return result
}

func (f *NodeFetcher) excluded(client clientInfo, exclude []server.NodeInfo) bool {
	for _, ex := range exclude {
		if ex.Client == client.client {
			return true
		}
	}
	return false
}

func (f *NodeFetcher) fetchID(c clientInfo, attempt int) (uint64, bool) {
	if attempt >= 5 {
		log.Printf("Unable to fetch ID for %s", c.URI)
		return 0, false
	}

	resp, err := c.client.Status(context.Background(), new(intra.StatusRequest))
	if err != nil {
		time.Sleep(5 * time.Second)
		return f.fetchID(c, attempt+1)
	}
	return resp.Id, true
}
