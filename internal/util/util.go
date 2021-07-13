package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ProcessResponse processes the response to an http request and unmarshalls it into the passed interface
func ProcessResponse(resp *http.Response, result interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error processing response: %v", err.Error())
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return fmt.Errorf("error processing response: %v", err.Error())
	}

	return resp.Body.Close()
}

// FormatUrl formats a url according to the passed format
func FormatUrl(baseUrl, format, endpoint string, params ...string) string {
	return fmt.Sprintf("%s%s", baseUrl, fmt.Sprintf(format, endpoint, params))
}
