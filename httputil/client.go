package httputil

import (
	"net"
	"net/http"
	"time"
)

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
		client.Transport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 10 * time.Second,
			IdleConnTimeout:       60 * time.Second,
			MaxIdleConns:          5,
			MaxIdleConnsPerHost:   5,
			DisableKeepAlives:     false,
		}
	}

	return &Client{client}
}
