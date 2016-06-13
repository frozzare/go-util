package httputil

import (
	"encoding/json"
	"net/http"
)

// GetJSON will bind JSON response from a url to a struct or return a error.
func GetJSON(v interface{}, target interface{}) error {
	config, err := getConfig(v)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", config.URL, nil)
	if err != nil {
		return err
	}

	res, err := config.Client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}
