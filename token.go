package shiptheory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ShiptheoryToken struct {
	accessToken string
	lastRefresh time.Time
}

type TokenResponseBody struct {

}

func (token *ShiptheoryToken) getAccessToken(username string, password string) string {
	body := map[string]string{
		"email": username,
		"password": password,
	}

	json_data, err := json.Marshal(body)
	fmt.Println(string(json_data))
    checkError(err)

	req, err := http.NewRequest(http.MethodPost, BASE_URL + "token", bytes.NewBuffer(json_data))
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

	fmt.Println(i)

	return "this_is_a_token"
}
