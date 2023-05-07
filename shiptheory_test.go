package shiptheory

import (
	"os"
	"testing"
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

	client.viewShipment(123)
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

	client.viewShipment(123)
}
