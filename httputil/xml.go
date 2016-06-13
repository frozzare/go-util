package httputil

import (
	"encoding/xml"
	"net/http"
)

// GetXML will bind XML response from a url to a struct or return a error.
func GetXML(v interface{}, target interface{}) error {
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

	return xml.NewDecoder(res.Body).Decode(target)
}
