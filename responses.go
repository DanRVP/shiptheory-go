package shiptheory

type TokenResponseBody struct {
	Success bool `json:"string"`
	Data TokenResponseBodyData `json:"data"`
}

type TokenResponseBodyData struct {
	Token string `json:"token"`
}

type ViewShipmentResponseBody struct {
	Shipment ViewShipmentResponseBodyShipment `json:"shipment"`
	Messages []ViewShipmentResponseBodyMessage `json:"messages"`
	Label string `json:"label"`
}

type ViewShipmentResponseBodyShipment struct {
	ChannelReferenceId string `json:"channel_reference_id"`
	ChannelReferenceId2 string `json:"channel_reference_id_2"`
	Created string `json:"created"`
	Modified string `json:"modified"`
	Status string `json:"status"`
	ChannelReferenceId3 string `json:"channel_reference_id_3"`
	DutyTaxNumber string `json:"duty_tax_number"`
	DutyTaxNumberType string `json:"duty_tax_number_type"`
	TrackingNumber string `json:"tracking_number"`
	Carrier string `json:"carrier"`
	ShippingPrice float64 `json:"shipping_price"`
	LatestPublicReference string `json:"latest_public_reference"`
	IsEuImport bool `json:"is_eu_import"`
	EligibleIossNumber string `json:"eligible_ioss_number"`
	IossNumber string `json:"ioss_number"`
	IossValueThreshold string `json:"ioss_value_threshold"`
	IsAustralianImport bool `json:"is_australian_import"`
	EligibleArn string `json:"eligible_arn"`
	EligibleAbn string `json:"eligible_abn"`
	TotalProductVolume float64 `json:"total_product_volume"`
	IsInternational bool `json:"is_international"`
}

type ViewShipmentResponseBodyMessage struct {
	Message string `json:"message"`
	Created string `json:"created"`
	Type string `json:"type"`
}

type BookShipmentResponseBody struct {
	Success bool `json:"success"`
	Status string `json:"status"`
	CarrierResult *BookShipmentResponseBodyCarrierResult `json:"carrier_result"`
	Message string `json:"message"`
}

type BookShipmentResponseBodyCarrierResult struct {
	Status string `json:"status"`
	Carrier string `json:"carrier"`
	ServiceName string `json:"service_name"`
	ServiceCode string `json:"service_code"`
	Tracking string `json:"tracking"`
	NoTracking *string `json:"no_tracking"`
	ShipmentCost *float64 `json:"shipment_cost"`
	Paperless bool `json:"paperless"`
}
