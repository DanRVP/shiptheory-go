package shiptheory

type TokenRequestBody struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type BookShipmentRequestBody struct {
	Reference string `json:"reference"`
	Reference2 string `json:"reference2"`
	DeliveryService int `json:"delivery_service,omitempty"`
	Increment int `json:"increment,omitempty"`
	ShipmentDetail BookShipmentRequestBodyShipmentDetail `json:"shipment_detail"`
	Recipient BookShipmentRequestBodyRecipient `json:"recipient"`
	Sender BookShipmentRequestBodySender `json:"sender"`
	Products []BookShipmentRequestBodyProduct `json:"products"`
	Packages []BookShipmentRequestBodyPackage `json:"packages"`
}

type BookShipmentRequestBodyShipmentDetail struct {
	Weight float64 `json:"weight"`
	Parcels int `json:"parcels"`
	Value float32 `json:"value"`
	ShippingPrice float32 `json:"shipping_price,omitempty"`
	Reference3 string `json:"reference3,omitempty"`
	EnhancementId int `json:"enhancement_id,omitempty"`
	FormatId int `json:"format_id,omitempty"`
	Instructions string `json:"instructions,omitempty"`
	GiftMessage string `json:"gift_message,omitempty"`
	ChannelShipserviceName string `json:"channel_shipservice_name,omitempty"`
	CurrencyCode string `json:"currency_code,omitempty"`
	SalesSource string `json:"sales_source,omitempty"`
	ShipDate string `json:"ship_date,omitempty"`
	RulesMetadata string `json:"rules_metadata,omitempty"`
	DutyTaxNumber string `json:"duty_tax_number,omitempty"`
	DutyTaxNumberType string `json:"duty_tax_number_type,omitempty"`
}

type BookShipmentRequestBodyAddress struct {
	Company string `json:"company"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2,omitempty"`
	AddressLine3 string `json:"address_line_3,omitempty"`
	City string `json:"city"`
	County string `json:"county,omitempty"`
	Country string `json:"country"`
	Postcode string `json:"postcode"`
	Telephone string `json:"telephone,omitempty"`
	Mobile string `json:"mobile,omitempty"`
	Email string `json:"email,omitempty"`
	TaxNumber string `json:"tax_number,omitempty"`
}

type BookShipmentRequestBodySender struct {
	BookShipmentRequestBodyAddress
}

type BookShipmentRequestBodyRecipient struct {
	BookShipmentRequestBodyAddress
	What3words string `json:"what3words,omitempty"`
	TaxNumbers []BookShipmentRequestBodyRecipientTaxNumber `json:"tax_numbers"`
}

type BookShipmentRequestBodyRecipientTaxNumber struct {
	TaxNumber string `json:"tax_number"`
	TaxNumberType string `json:"tax_number_type"`
}

type BookShipmentRequestBodyProduct struct {
	Name string `json:"name"`
	Sku string `json:"sku"`
	Qty int `json:"qty,omitempty"`
	Value float64 `json:"value"`
	Weight float64 `json:"weight"`
	CommodityCode string `json:"commodity_code,omitempty"`
	CommodityDescription string `json:"commodity_description,omitempty"`
	CommodityManucountry string `json:"commodity_manucountry,omitempty"`
}

type BookShipmentRequestBodyPackage struct {
	Id string `json:"id"`
	Weight string `json:"weight"`
}
