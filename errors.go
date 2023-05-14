package shiptheory

type TokenErrorBody struct {
	Message string `json:"message"`
	Url string `json:"url"`
	Code int `json:"code"`
}

type ShipmentErrorBody struct {
	Message string `json:"message,omitempty"`
	Url string `json:"url,omitempty"`
	Code int `json:"code,omitempty"`
	Error interface{} `json:"error,omitempty"`
}
