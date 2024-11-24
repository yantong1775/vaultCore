package token

import "time"

// Maker is a interface that creates and validates tokens
type Maker interface {
	// CreateToken create a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)
	// Check if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
