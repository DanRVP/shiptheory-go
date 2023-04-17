package shiptheory

import (
	"testing"
)


// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestOutputMessage(t *testing.T) {
    outputMsg("test")
}

func TestToken(t *testing.T) {
	client := ShiptheoryClient{
		username: "",
		password: "",
	}

	client.validateToken()
}
