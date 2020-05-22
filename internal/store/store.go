package store

import (
	"context"
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Storage level errors
var (
	ErrNotFound = errors.New("Not found")
)

// Paste represents a data to be pasted
type Paste struct {
	Title     string
	ID        uuid.UUID
	Body      string
	Headers   map[string]string
	CreatedAt time.Time
	Expiry    time.Time
}

// Store defines the interface for paste storage
type Store interface {
	SavePaste(context.Context, Paste) (Paste, error)
	GetPaste(ctx context.Context, id uuid.UUID) (Paste, error)
}
