package shiptheory

import (
	"log"
	"time"
)

const BASE_URL = "https://api.shiptheory.com/v1/"

type ShiptheoryClient struct {
	token ShiptheoryToken
	username string
	password string
}


func (client ShiptheoryClient) validateToken() {
	if len(client.token.accessToken) < 1 || client.checkTokenLifeExpired() {
		client.token.getAccessToken(client.username, client.password)
	}
}

func (client ShiptheoryClient) checkTokenLifeExpired() bool {
	now := time.Now()
	diff := now.Sub(client.token.lastRefresh)
	return diff.Minutes() > 58 // Refresh 2 mins before token expires. Tokens last for 1 hr.
}

func checkError(err error) {
	if err != nil {
        log.Fatal(err)
    }
}