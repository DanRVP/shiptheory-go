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

// Make a GET request to `shipments/view/{ref}`
func (client *ShiptheoryClient) ViewShipment(id int) (res_body ViewShipmentResponseBody, err error) {
	err = client.validateToken(client.username, client.password)
	checkError(err)

	endpoint := "shipments/" + strconv.Itoa(id)
	var err_body ShipmentErrorBody = ShipmentErrorBody{}
	err = client.makeShiptheoryApiRequest(http.MethodGet, endpoint, nil, &res_body, &err_body)
	checkError(err)

	if err_body.Message != "" {
		return res_body, errors.New(err_body.Message)
	}

	return res_body, nil
}

// Make a POST request to `shipments`
func (client *ShiptheoryClient) BookShipment(body BookShipmentRequestBody) (res_body BookShipmentResponseBody, err error) {
	err = client.validateToken(client.username, client.password)
	checkError(err)

	var err_body TokenErrorBody = TokenErrorBody{}
	err = client.makeShiptheoryApiRequest(http.MethodPost, "shipments", body, &res_body, &err_body)
	checkError(err)

	if err_body.Message != "" {
		return res_body, errors.New(err_body.Message)
	}

	return res_body, nil
}

// Make a GET request to `shipments/list`
// func (client *ShiptheoryClient) ListShipments(params map[string][string]) (res_body ViewShipmentResponseBody, err error) {
// 	err = client.validateToken(client.username, client.password)
// 	checkError(err)

	// endpoint := "shipments/list?" + buildQueryString(params)
// 	var err_body ShipmentErrorBody = ShipmentErrorBody{}
// 	err = client.makeShiptheoryApiRequest(http.MethodGet, endpoint, nil, &res_body, &err_body)
// 	checkError(err)

// 	if err_body.Message != "" {
// 		return res_body, errors.New(err_body.Message)
// 	}

// 	return res_body, nil
// }
