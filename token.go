package shiptheory

import (
	"errors"
	"net/http"
	"time"
)

type ShiptheoryToken struct {
	accessToken string
	lastRefresh time.Time
}

func (token *ShiptheoryToken) validateToken(username string, password string) error {
	if len(token.accessToken) < 1 || token.checkTokenLifeExpired() {
		return token.refreshAccessToken(username, password)
	}

	return nil
}

func (token *ShiptheoryToken) checkTokenLifeExpired() bool {
	now := time.Now()
	diff := now.Sub(token.lastRefresh)
	return diff.Minutes() > 58 // Refresh 2 mins before token expires. Tokens last for 1 hr.
}

func (token *ShiptheoryToken) refreshAccessToken(username string, password string) (error) {
	var body TokenRequestBody = TokenRequestBody{
		Email: username,
		Password: password,
	}

	var res_body TokenResponseBody = TokenResponseBody{}
	var err_body TokenErrorBody = TokenErrorBody{}
	err := makeShiptheoryApiRequest(http.MethodPost, "token", body, &res_body, &err_body)
	checkError(err)

	if err_body.Message != "" {
		return errors.New(err_body.Message)
	} else {
		token.accessToken = res_body.Data.Token
		token.lastRefresh = time.Now()

		return nil
	}
}
