package shiptheory

import (
	"net/http"
	"time"
)

type ShiptheoryToken struct {
	accessToken string
	lastRefresh time.Time
}

type TokenResponseBody struct {

}
