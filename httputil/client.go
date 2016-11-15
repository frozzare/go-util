package httputil

import "net/http"

// Client represents the custom http client.
type Client struct {
	*http.Client
}

// NewClient will create a new http client instance.
func NewClient(client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	return &Client{client}
}
