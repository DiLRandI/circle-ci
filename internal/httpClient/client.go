package httpclient

import (
	"fmt"
	"net/http"
	"time"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	httpClient *http.Client
}

func New(timeout time.Duration) Client {
	return &client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do the request: %w", err)
	}

	return resp, nil
}
