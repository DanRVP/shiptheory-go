package shiptheory

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
	"errors"
	"io"
)

// Shiptheory base URL.
const BASE_URL = "https://api.shiptheory.com/v1/"

// Make a request to Shiptheory and process the result.
func (shiptheoryClient *ShiptheoryClient) makeShiptheoryApiRequest(method string, endpoint string, body interface{}, success_body interface{}, error_body interface{}) (error) {
	var json_data []byte
	var err error

	if !isValidHttpMethod(method) {
		return errors.New("invalid HTTP method not supported by Shiptheory API")
	}

	var req *http.Request
	if body != nil {
		json_data, err = json.Marshal(body)
		if err != nil {
			return err
		}

		req, err = http.NewRequest(method, BASE_URL + endpoint, bytes.NewBuffer(json_data))
	} else {
		req, err = http.NewRequest(method, BASE_URL + endpoint, nil)
	}

	if err != nil {
		return err
	}

	shiptheoryClient.appendHeaders(req)

	client := http.Client{
		Timeout: time.Duration(5) * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	res_body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		err = json.Unmarshal(res_body, &success_body)
	} else {
		err = json.Unmarshal(res_body, &error_body)
	}

	res.Body.Close()
	return err
}

// Validate whether a string is a valid HTTP method supported by Shiptheory's API.
func isValidHttpMethod(method string) bool {
	valid_methods := [3]string{http.MethodGet, http.MethodPost, http.MethodPut}
	for _, http_method := range valid_methods {
		if http_method == method {
			return true
		}
	}

	return false
}

// Append appropriate headers for a request to Shiptheory to a http.Request object.
func (shiptheoryClient *ShiptheoryClient) appendHeaders(req *http.Request) {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if shiptheoryClient.token.accessToken != "" {
		req.Header.Add("Authorization", "Bearer " + shiptheoryClient.token.accessToken)
	}

	if shiptheoryClient.partnerTag != "" {
		req.Header.Add("Shiptheory-Partner-Tag", shiptheoryClient.partnerTag)
	}
}

func buildQueryString(params map[string]string) string {
	query := url.Values{}
	for key, value := range params {
		query.Add(key, value)
	}

	return query.Encode()
}
