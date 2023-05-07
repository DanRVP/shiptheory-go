package shiptheory

import (
	"net/http"
	"strconv"
)

type ShiptheoryClient struct {
	token ShiptheoryToken
	username string
	password string
	partner_tag string
}

func (client *ShiptheoryClient) viewShipment(id int) (shipment string, err error) {
	err = client.token.validateToken(client.username, client.password)
	checkError(err)

	endpoint := "/shipments/" + strconv.Itoa(id)
	var res_body TokenResponseBody = TokenResponseBody{}
	var err_body TokenErrorBody = TokenErrorBody{}
	err = makeShiptheoryApiRequest(http.MethodGet, endpoint, nil, &res_body, &err_body)

	return shipment, err
}
