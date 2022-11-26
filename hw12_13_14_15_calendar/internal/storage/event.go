package storage

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID               uuid.UUID `db:"id"`
	Title            string    `db:"title"`
	BeginTime        time.Time `db:"begin_time"`
	EndTime          time.Time `db:"end_time"`
	Description      string    `db:"description"`
	UserID           uuid.UUID `db:"user_id"`
	NotificationTime time.Time `db:"notification_time"`
}

func (e *Event) Duration() time.Duration {
	return e.EndTime.Sub(e.BeginTime)
}

func (e *Event) inPeriod(from time.Time, to time.Time) bool {
	return e.BeginTime.Before(to) && !e.EndTime.Before(from)
}
