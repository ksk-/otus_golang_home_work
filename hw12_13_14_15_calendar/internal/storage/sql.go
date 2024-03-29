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
	getEventToNotifyQuery = `
SELECT id, title, begin_time, end_time, coalesce(description, '') as description, user_id, notification_time 
FROM events 
WHERE notification_time >= $1 AND notification_time < $2
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
	deleteEventQuery      = `DELETE FROM events WHERE id = $1`
	deletePastEventsQuery = `DELETE FROM events WHERE end_time < $1`
)

func NewSQLStorage(db *sqlx.DB, logger *logger.Logger) Storage {
	return &sqlStorage{db: db, logger: logger}
}

type sqlStorage struct {
	db     *sqlx.DB
	logger *logger.Logger
}

func (s *sqlStorage) Close() error {
	return s.db.Close()
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
		return nil, fmt.Errorf("%w: %v", ErrEventNotFound, err)
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

func (s *sqlStorage) GetEventsToNotify(ctx context.Context, from, to time.Time) ([]Event, error) {
	events := make([]Event, 0)
	if err := s.db.SelectContext(ctx, &events, getEventToNotifyQuery, from, to); err != nil {
		return nil, fmt.Errorf("get events to notify: %w", err)
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
	r, err := s.db.NamedExecContext(ctx, updateEventQuery, event)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	updated, err := r.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if updated == 0 {
		return ErrEventNotFound
	}

	return nil
}

func (s *sqlStorage) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	r, err := s.db.ExecContext(ctx, deleteEventQuery, id)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	deleted, err := r.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if deleted == 0 {
		return ErrEventNotFound
	}

	return nil
}

func (s *sqlStorage) DeletePastEvents(ctx context.Context, before time.Time) (int64, error) {
	r, err := s.db.ExecContext(ctx, deletePastEventsQuery, before)
	if err != nil {
		return 0, fmt.Errorf("delete past event: %w", err)
	}
	return r.RowsAffected()
}
