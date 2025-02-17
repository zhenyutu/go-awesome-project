package geecache

import (
	"awesomeProject/project/geecache/consistenthash"
	"errors"
	"net/http"
	"strings"
	"sync"
)

const (
	defaultBasePath = "/_geecache/"
	defaultReplicas = 50
)

type CacheHttpServer struct {
	self     string
	basePath string
	mux      sync.Mutex
	peers    *consistenthash.HashRing
	servers  map[string]*CacheHttpServer
}

func (g *Group) newHttpServer(self string) *CacheHttpServer {
	return &CacheHttpServer{
		self:     self,
		basePath: defaultBasePath,
	}
}

func (g *Group) registerHttpServer(server *CacheHttpServer) {
	server.mux.Lock()
	defer server.mux.Unlock()

	if server.peers == nil {
		server.peers = consistenthash.New(defaultReplicas, nil)
	}
	server.peers.Add(server.self)
	server.servers[server.self] = server
}

/**
 * Pick Server
 */
func (server *CacheHttpServer) Pick(key string) (PeerGetter, error) {
	server.mux.Lock()
	defer server.mux.Unlock()

	if server.peers == nil {
		return nil, errors.New("peer not exists")
	}
	if peer := server.peers.Get(key); peer != "" && peer != server.self {
		return server.servers[peer], nil
	}
	return nil, errors.New("peer not exists")
}

func (server *CacheHttpServer) PeerGet(group string, key string) ([]byte, error) {
	cacheClient := HttpCacheClient{server.self}
	return cacheClient.Get(group, key)
}

/**
 * Server start
 */
func (s *CacheHttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parts := strings.SplitN(r.URL.Path[len(s.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := parts[0]
	key := parts[1]

	group := getGroup(groupName)
	if group == nil {
		http.Error(w, "group not found", http.StatusNotFound)
	}

	data, ok := group.Get(key)
	if ok {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(data.data)
		return
	}

	http.Error(w, "not found", http.StatusInternalServerError)
}
