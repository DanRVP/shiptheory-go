package shiptheory

import (
	"time"
)

type ShiptheoryToken struct {
	accessToken string
	lastRefresh time.Time
}

// Check that the current token is still 'alive'
func (token *ShiptheoryToken) checkTokenLifeExpired() bool {
	now := time.Now()
	diff := now.Sub(token.lastRefresh)
	return diff.Minutes() > 58 // Refresh 2 mins before token expires. Tokens last for 1 hr.
}
