package geecache

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

/**
 * HTTP
 */
type HttpCacheClient struct {
	baseURL string
}

func (server *HttpCacheClient) Get(group string, key string) ([]byte, error) {
	u := fmt.Sprintf(
		"%v%v/%v",
		server.baseURL,
		url.QueryEscape(group),
		url.QueryEscape(key),
	)

	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get: %v", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
