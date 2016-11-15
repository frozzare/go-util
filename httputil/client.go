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

	// Set a custom default transport instead of the default transport.
	// https://golang.org/src/net/http/transport.go#L33
	if client.Transport == nil {
		client.Transport = DefaultTransport
	}

	return &Client{client}
}
