package auth

import (
	"time"

	"github.com/skerkour/golibs/uuid"
)

type ApiKey struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time

	Name string
	// ExpiresAt *time.Time
}
