package httputil

import (
	"errors"
	"net/http"
	"reflect"
	"time"
)

// Config represents httputil config.
type Config struct {
	Dualstack  bool
	Client     *http.Client
	SkipVerify bool
	Timeout    time.Duration
	URL        string
}

// getConfig will create a Config instance of a interface.
func getConfig(v interface{}) (*Config, error) {
	if reflect.ValueOf(v).Kind() == reflect.String {
		v = &Config{URL: v.(string)}
	}

	config, ok := v.(*Config)
	if !ok {
		return nil, errors.New("GetJSON requires httputil.Config struct")
	}

	if config.Client == nil {
		config.Client = NewClient(config.Timeout, true, true)
	}

	return config, nil
}
