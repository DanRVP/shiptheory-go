package shiptheory

import (
	"time"
	"http"
)

type ShiptheoryClient struct {
	token ShiptheoryToken
	username string
	password string
}

func (client ShiptheoryClient) validateToken() {
	var err error
	if len(client.token.accessToken) < 1 || client.checkTokenLifeExpired() {
		client.token, err = client.getAccessToken(client.username, client.password)
		checkError(err)
	}
}

func (client ShiptheoryClient) checkTokenLifeExpired() bool {
	now := time.Now()
	diff := now.Sub(client.token.lastRefresh)
	return diff.Minutes() > 58 // Refresh 2 mins before token expires. Tokens last for 1 hr.
}

func (client ShiptheoryClient) getAccessToken(username string, password string) (ShiptheoryToken, error) {
	body := map[string]string{
		"email": username,
		"password": password,
	}

	token_response, err := makeShiptheoryApiRequest(http.MethodPost, "token", body)
	checkError(err)
	client.token.accessToken = token_response.data.token
}
