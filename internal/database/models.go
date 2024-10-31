package database

import (
	"time"
	"github.com/google/uuid"
)

// User represents one person in our system
// Why these fields?
// - UUID: Better than auto-increment for distributed systems (no central counter needed)
// - Timestamps: Always need to know when things happened
// - Name: The simplest way to identify users (no email/password yet - YAGNI)
type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}
