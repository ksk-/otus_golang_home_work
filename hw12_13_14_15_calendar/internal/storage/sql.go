package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/logger"
)

const (
	listEventsQuery = `
SELECT id, title, begin_time, end_time, coalesce(description, '') as description, user_id, notification_time 
FROM events
ORDER BY begin_time
LIMIT $1 OFFSET $2
`
	getEventQuery = `
SELECT id, title, begin_time, end_time, coalesce(description, '') as description, user_id, notification_time 
FROM events 
WHERE id = $1
`
	getEventsForPeriodQuery = `
SELECT id, title, begin_time, end_time, coalesce(description, '') as description, user_id, notification_time 
FROM events 
WHERE begin_time < $2 AND end_time >= $1
ORDER BY begin_time
`
	insertEventQuery = `
INSERT INTO events 
VALUES (:id, :title, :begin_time, :end_time, :description, :user_id, :notification_time)
`
	updateEventQuery = `
UPDATE events 
SET 
    title=:title, 
    begin_time=:begin_time, 
    end_time=:end_time, 
    description=nullif(:description,''), 
    user_id=:user_id, 
    notification_time=:notification_time
WHERE id = :id
`
	deleteEventQuery = `DELETE FROM events WHERE id = $1`
)

func NewSQLStorage(db *sqlx.DB, logger *logger.Logger) Storage {
	return &sqlStorage{db: db, logger: logger}
}

type sqlStorage struct {
	db     *sqlx.DB
	logger *logger.Logger
}

func (s *sqlStorage) ListEvents(ctx context.Context, limit, offset uint64) ([]Event, error) {
	events := make([]Event, 0)
	if err := s.db.SelectContext(ctx, &events, listEventsQuery, limit, offset); err != nil {
		return nil, fmt.Errorf("get events: %w", err)
	}
	return events, nil
}

func (s *sqlStorage) GetEvent(ctx context.Context, id uuid.UUID) (*Event, error) {
	var event Event
	if err := s.db.GetContext(ctx, &event, getEventQuery, id); err != nil {
		return nil, fmt.Errorf("get event: %w", err)
	}
	return &event, nil
}

func (s *sqlStorage) GetEventsForPeriod(ctx context.Context, from, to time.Time) ([]Event, error) {
	events := make([]Event, 0)
	if err := s.db.SelectContext(ctx, &events, getEventsForPeriodQuery, from, to); err != nil {
		return nil, fmt.Errorf("get events for period: %w", err)
	}
	return events, nil
}

func (s *sqlStorage) InsertEvent(ctx context.Context, event *Event) error {
	if _, err := s.db.NamedExecContext(ctx, insertEventQuery, event); err != nil {
		return fmt.Errorf("insert event: %w", err)
	}
	return nil
}

func (s *sqlStorage) UpdateEvent(ctx context.Context, event *Event) error {
	if _, err := s.db.NamedExecContext(ctx, updateEventQuery, event); err != nil {
		return fmt.Errorf("update event: %w", err)
	}
	return nil
}

func (s *sqlStorage) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	if _, err := s.db.ExecContext(ctx, deleteEventQuery, id); err != nil {
		return fmt.Errorf("delete event: %w", err)
	}
	return nil
}
