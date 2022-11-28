package storage

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEventNotFound      = errors.New("event not found")
	ErrEventAlreadyExists = errors.New("event already exists")
)

type Storage interface {
	// ListEvents list stored events sorted by begin time
	ListEvents(ctx context.Context, limit, offset uint64) ([]Event, error)

	// GetEvent returns the specified event if it exists
	GetEvent(ctx context.Context, id uuid.UUID) (*Event, error)

	// GetEventsForPeriod return events for the specified period `[from,to)` sorted by begin time
	GetEventsForPeriod(ctx context.Context, from, to time.Time) ([]Event, error)

	// InsertEvent insert the specified event
	InsertEvent(ctx context.Context, event *Event) error

	// UpdateEvent insert the specified event if it exists
	UpdateEvent(ctx context.Context, event *Event) error

	// DeleteEvent delete the specified event if it exists
	DeleteEvent(ctx context.Context, id uuid.UUID) error
}
