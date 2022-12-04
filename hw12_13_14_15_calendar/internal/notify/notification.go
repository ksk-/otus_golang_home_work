package notify

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uuid.UUID `json:"id"`
	EventID   uuid.UUID `json:"eventId"`
	Title     string    `json:"title"`
	EventTime time.Time `json:"eventTime"`
	UserID    uuid.UUID `json:"userId"`
}
