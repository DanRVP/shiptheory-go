package shiptheory

import (
	"os"
	"testing"
	"fmt"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestOutputMessage(t *testing.T) {
    outputMsg("test")
}

// Test getting a token from Shiptheory.
func TestToken(t *testing.T) {
	pw, err := os.ReadFile("./password.txt")
	checkError(err)
	password := string(pw)

	client := ShiptheoryClient{
		username: "dan.rogers@shiptheory.com",
		password: password,
	}

	client.validateToken()
	if client.token != nil {
		fmt.Println(client.token.accessToken)
	}
}