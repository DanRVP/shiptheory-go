package shiptheory

type TokenErrorBody struct {
	Message string `json:"message"`
	Url string `json:"url"`
	Code int `json:"code"`
}