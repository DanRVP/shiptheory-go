package shiptheory

import (
	"net/http"
	"time"
	"log"
	"encoding/json"
	"bytes"
)

type ShiptheoryToken struct {
	accessToken string
	lastRefresh time.Time
}

type TokenResponseBody struct {

}

func (token *ShiptheoryToken) getAccessToken(username string, password string) bool {
	body := map[string]string{
		"username": username,
		"password": password,
	}

	json_data, err := json.Marshal(body)
    if err != nil {
        log.Fatal(err)
    }

	req, err := http.NewRequest(http.MethodPost, BASE_URL + "/token", bytes.NewBuffer(json_data))
    if err != nil {
        log.Fatal(err)
    }

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{
		Timeout: time.Duration(5) * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
        log.Fatal(err)
    }


}
