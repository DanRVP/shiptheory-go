package shiptheory

import (
	"os"
	"testing"
	"time"
	"strconv"
)

// Test getting a token from Shiptheory.
func TestToken(t *testing.T) {
	pw, err := os.ReadFile("./password.txt")
	checkError(err)
	password := string(pw)

	client := ShiptheoryClient{
		username: "dan.rogers@shiptheory.com",
		password: password,
	}

	client.ViewShipment(123)
}

func TestViewShipment(t *testing.T) {
	pw, err := os.ReadFile("./password.txt")
	checkError(err)
	password := string(pw)

	client := ShiptheoryClient{
		username: "dan.rogers@shiptheory.com",
		password: password,
		token: ShiptheoryToken{},
	}

	_, err = client.ViewShipment(680)
	checkError(err)
}

func TestBookShipment(t *testing.T) {
	pw, err := os.ReadFile("./password.txt")
	checkError(err)
	password := string(pw)

	client := ShiptheoryClient{
		username: "dan.rogers@shiptheory.com",
		password: password,
		token: ShiptheoryToken{},
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	body := BookShipmentRequestBody{
		Reference: "S-" + timestamp,
		Reference2: "O-" + timestamp,
		ShipmentDetail: BookShipmentRequestBodyShipmentDetail{
			Weight: 1,
			Parcels: 1,
			Value: 5,
			ShippingPrice: 2.99,
			Reference3: "ORDERREF3",
			SalesSource: "Shiptheory Go Library",
			RulesMetadata: "my_custom_string",
			ChannelShipserviceName: "Next Day",
			CurrencyCode: "GBP",
		},
		Recipient: BookShipmentRequestBodyRecipient{
			BookShipmentRequestBodyAddress: BookShipmentRequestBodyAddress{
				Company: "Shiptheory",
				Firstname: "Test",
				Lastname: "Customer",
				AddressLine1: "Unit 4.1 Paintworks",
				AddressLine2: "Bath Road",
				City: "Bristol",
				Postcode: "BS4 3EH",
				Email: "recipient@test.com",
				Telephone: "01234567890",
			},
			What3words: "///what.three.words",
			TaxNumbers: []BookShipmentRequestBodyRecipientTaxNumber{
				{
					TaxNumber: "GB205672212000",
					TaxNumberType: "EORI",
				},
				{
					TaxNumber: "GB123456789",
					TaxNumberType: "VAT",
				},
			},
		},
		Sender: BookShipmentRequestBodySender{
			BookShipmentRequestBodyAddress: BookShipmentRequestBodyAddress{
				Company: "Shiptheory",
				Firstname: "Test",
				Lastname: "Shipper",
				AddressLine1: "Bristol Old Vic",
				AddressLine2: "King Street",
				City: "Bristol",
				Postcode: "BS1 4ED",
				Email: "sender@test.com",
				Telephone: "01234567890",
			},
		},
		Products: []BookShipmentRequestBodyProduct{
			{
				Name: "Test Product",
				Sku: "TestProduct1",
				Qty: 5,
				Value: 15.99,
				Weight: 2.67,
				CommodityCode: "8443991000",
				CommodityManucountry: "GB",
			},
			{
				Name: "Test Product 2",
				Sku: "TestProduct2",
				Qty: 3,
				Value: 6.75,
				Weight: 0.84,
				CommodityCode: "8443991000",
				CommodityManucountry: "GB",
			},
		},
	}

	_, err = client.BookShipment(body)
	checkError(err)
}

// func TestVListShipments(t *testing.T) {
// 	pw, err := os.ReadFile("./password.txt")
// 	checkError(err)
// 	password := string(pw)

// 	client := ShiptheoryClient{
// 		username: "dan.rogers@shiptheory.com",
// 		password: password,
// 		token: ShiptheoryToken{},
// 	}

// 	params := make(map[string]string)
// 	params["created_from"] = "2022-05-19"

// 	_, err = client.ListShipments(680)
// 	checkError(err)
// }

func TestCamelToSnake(t *testing.T) {
	t.Log(camelToSnake("ChannelReferenceId2"))
}

func TestBuildQueryString(t *testing.T) {
	query := ListShipmentsQuery {
		ChannelName: "Magento 2",
	}

	res := buildQueryString(query)
	t.Log(res)
}
