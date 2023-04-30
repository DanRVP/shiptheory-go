package shiptheory

import (
	"errors"
	"net/http"
	"time"
)

type ShiptheoryClient struct {
	token *ShiptheoryToken
	username string
	password string
}

func (client *ShiptheoryClient) validateToken() {
	var err error
	if client.token == nil || len(client.token.accessToken) < 1 || client.checkTokenLifeExpired() {
		client.token, err = client.getAccessToken(client.username, client.password)
		checkError(err)
	}
}

func (client *ShiptheoryClient) checkTokenLifeExpired() bool {
	now := time.Now()
	diff := now.Sub(client.token.lastRefresh)
	return diff.Minutes() > 58 // Refresh 2 mins before token expires. Tokens last for 1 hr.
}

func (client *ShiptheoryClient) getAccessToken(username string, password string) (*ShiptheoryToken, error) {
	var body TokenRequestBody = TokenRequestBody{
		Email: username,
		Password: password,
	}

	var res_body TokenResponseBody = TokenResponseBody{}
	var err_body TokenErrorBody = TokenErrorBody{}
	err := makeShiptheoryApiRequest(http.MethodPost, "token", body, &res_body, &err_body)
	checkError(err)

	if err_body.Message != "" {
		return nil, errors.New(err_body.Message)
	} else {
		var token ShiptheoryToken = ShiptheoryToken{
			accessToken: res_body.Data.Token,
			lastRefresh: time.Now(),
		}

		return &token, nil
	}

}
