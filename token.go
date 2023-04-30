package shiptheory

import (
	"time"
)

type ShiptheoryToken struct {
	accessToken string
	lastRefresh time.Time
}

type TokenResponseBody struct {
	Success bool `json:"string"`
	Data TokenResponseBodyPartData `json:"data"`
}

type TokenResponseBodyPartData struct {
	Token string `json:"token"`
}

type TokenRequestBody struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type TokenErrorBody struct {
	Message string `json:"message"`
	Url string `json:"url"`
	Code int `json:"code"`
}
