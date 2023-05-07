package shiptheory

import (
	"net/http"
	"strconv"
	"time"
	"errors"
)

type ShiptheoryClient struct {
	token ShiptheoryToken
	username string
	password string
	partnerTag string
}

// Check if the client still has a valid access token. If it doesn't, then it gets a new one.
func (client *ShiptheoryClient) validateToken(username string, password string) error {
	if len(client.token.accessToken) < 1 || client.token.checkTokenLifeExpired() {
		return client.refreshAccessToken(username, password)
	}

	return nil
}

// Make a request to Shiptheory to get a new access token and save it to the client.
func (client *ShiptheoryClient) refreshAccessToken(username string, password string) (error) {
	var body TokenRequestBody = TokenRequestBody{
		Email: username,
		Password: password,
	}

	var res_body TokenResponseBody = TokenResponseBody{}
	var err_body TokenErrorBody = TokenErrorBody{}
	err := client.makeShiptheoryApiRequest(http.MethodPost, "token", body, &res_body, &err_body)
	checkError(err)

	if err_body.Message != "" {
		return errors.New(err_body.Message)
	} else {
		client.token = ShiptheoryToken{
			accessToken: res_body.Data.Token,
			lastRefresh: time.Now(),
		}

		return nil
	}
}

// Make a request to `shipments/view/{ref}`
func (client *ShiptheoryClient) ViewShipment(id int) (shipment string, err error) {
	err = client.validateToken(client.username, client.password)
	checkError(err)

	endpoint := "shipments/" + strconv.Itoa(id)
	var res_body ViewShipmentResponseBody = ViewShipmentResponseBody{}
	var err_body TokenErrorBody = TokenErrorBody{}
	err = client.makeShiptheoryApiRequest(http.MethodGet, endpoint, nil, &res_body, &err_body)

	return shipment, err
}
