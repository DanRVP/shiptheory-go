package shiptheory

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	"strings"
	"errors"
	"log"
	"io"
)

const BASE_URL = "https://api.shiptheory.com/v1/"

func makeShiptheoryApiRequest(method string, endpoint string, body interface{}, success_body interface{}, error_body interface{}) (error) {
	var json_data []byte
	var err error
	var http_method string

	switch strings.ToUpper(method) {
		case http.MethodPost:
			http_method = http.MethodPost
		case http.MethodGet:
			http_method = http.MethodGet
		case http.MethodPut:
			http_method = http.MethodPut
		default:
			return errors.New("Invalid HTTP method not supported by Shiptheory API")
	}

	var req *http.Request
	if body != nil {
		json_data, err = json.Marshal(body)
		if err != nil {
			return err
		}

		req, err = http.NewRequest(http_method, BASE_URL + endpoint, bytes.NewBuffer(json_data))
	} else {
		req, err = http.NewRequest(http_method, BASE_URL + endpoint, nil)
	}

	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

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

func checkError(err error) {
	if err != nil {
        log.Fatal(err)
    }
}
