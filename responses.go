package shiptheory

type TokenResponseBody struct {
	Success bool `json:"string"`
	Data TokenResponseBodyPartData `json:"data"`
}

type TokenResponseBodyPartData struct {
	Token string `json:"token"`
}