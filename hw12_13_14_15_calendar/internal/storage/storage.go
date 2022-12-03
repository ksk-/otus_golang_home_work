package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/logger"
)

var (
	ErrEventNotFound      = errors.New("event not found")
	ErrEventAlreadyExists = errors.New("event already exists")
)

type Storage interface {
	io.Closer

	// ListEvents lists stored events sorted by begin time
	ListEvents(ctx context.Context, limit, offset uint64) ([]Event, error)

	// GetEvent returns the specified event if it exists
	GetEvent(ctx context.Context, id uuid.UUID) (*Event, error)

	// GetEventsForPeriod returns events for the specified period `[from,to)` sorted by begin time
	GetEventsForPeriod(ctx context.Context, from, to time.Time) ([]Event, error)

	// GetEventsToNotify returns events to be notified about during the specified period `[from,to)`
	GetEventsToNotify(ctx context.Context, from, to time.Time) ([]Event, error)

	// InsertEvent inserts the specified event
	InsertEvent(ctx context.Context, event *Event) error

	// UpdateEvent inserts the specified event if it exists
	UpdateEvent(ctx context.Context, event *Event) error

	// DeleteEvent deletes the specified event if it exists
	DeleteEvent(ctx context.Context, id uuid.UUID) error

	// DeletePastEvents delete events past before the specified time
	DeletePastEvents(ctx context.Context, before time.Time) (int64, error)
}

func NewStorage(cfg config.Storage, logger *logger.Logger) (Storage, error) {
	switch cfg.Type {
	case "memory":
		return NewMemoryStorage(), nil
	case "sql":
		db, err := sqlx.Connect("pgx", cfg.DB.DSN())
		if err != nil {
			return nil, fmt.Errorf("create DB connection: %w", err)
		}
		return NewSQLStorage(db, logger), nil
	default:
		return nil, fmt.Errorf("unknown storage type: %s", cfg.Type)
	}
}
