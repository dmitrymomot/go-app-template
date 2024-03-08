package dto

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user entity.
type User struct {
	ID        uuid.UUID
	Email     string
	CreatedAt time.Time
}
