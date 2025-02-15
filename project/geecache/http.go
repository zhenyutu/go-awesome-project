package geecache

import (
	"net/http"
	"strings"
)

const defaultBasePath = "/_geecache/"

type CacheHttpServer struct {
	self     string
	basePath string
}

func (g *Group) newHttpServer(self string) *CacheHttpServer {
	return &CacheHttpServer{
		self:     self,
		basePath: defaultBasePath,
	}
}

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
