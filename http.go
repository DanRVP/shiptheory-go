package shiptheory

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
	"strings"
	"errors"
	"log"
)

const BASE_URL = "https://api.shiptheory.com/v1/"

func makeShiptheoryApiRequest(method string, endpoint string, body interface{}) (interface{}, error) {
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
			return nil, errors.New("Invalid HTTP method not supported by Shiptheory API")
	}

	if body != nil {
		json_data, err = json.Marshal(body)
		checkError(err)
	}

	req, err := http.NewRequest(http_method, BASE_URL + endpoint, bytes.NewBuffer(json_data))
    checkError(err)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{
		Timeout: time.Duration(5) * time.Second,
	}

	res, err := client.Do(req)
	checkError(err)

	res_body, err := io.ReadAll(res.Body)
	checkError(err)

	var i interface{}
	err = json.Unmarshal(res_body, &i)
	res.Body.Close()
	checkError(err)

	return i, nil
}

func checkError(err error) {
	if err != nil {
        log.Fatal(err)
    }
}
