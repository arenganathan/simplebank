package token

import "time"

// Maker is a interface for token management
type Maker interface {
	// createed a new token for specific user with duration
	CreateToken(username string, duration time.Duration) (string, error)

	//
	VerifyToken(token string) (*Payload, error)
}
