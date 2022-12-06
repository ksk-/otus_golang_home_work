package app

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/logger"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/storage"
)

type CreateEventCommand struct {
	Title            string
	BeginTime        time.Time
	EndTime          time.Time
	Description      string
	UserID           uuid.UUID
	NotificationTime time.Time
}

type App struct {
	logger  *logger.Logger
	storage storage.Storage
}

func New(storage storage.Storage, logger *logger.Logger) *App {
	return &App{storage: storage, logger: logger}
}

func (a *App) CreateEvent(ctx context.Context, cmd *CreateEventCommand) (uuid.UUID, error) {
	eventID := uuid.New()
	event := storage.Event{
		ID:               eventID,
		Title:            cmd.Title,
		BeginTime:        cmd.BeginTime,
		EndTime:          cmd.EndTime,
		Description:      cmd.Description,
		UserID:           cmd.UserID,
		NotificationTime: cmd.NotificationTime,
	}

	if err := a.storage.InsertEvent(ctx, &event); err != nil {
		return uuid.Nil, fmt.Errorf("insert event to storage: %w", err)
	}
	a.logger.Debug(fmt.Sprintf("[eventID: %v]: event created", eventID))
	return eventID, nil
}

func (a *App) UpdateEvent(ctx context.Context, event *storage.Event) error {
	if err := a.storage.UpdateEvent(ctx, event); err != nil {
		return fmt.Errorf("update event in storage: %w", err)
	}
	a.logger.Debug(fmt.Sprintf("[eventID: %v]: event updated", event.ID))
	return nil
}

func (a *App) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	if err := a.storage.DeleteEvent(ctx, eventID); err != nil {
		return fmt.Errorf("delete event from storage: %w", err)
	}
	a.logger.Debug(fmt.Sprintf("[eventID: %v]: event deleted", eventID))
	return nil
}

func (a *App) GetEvent(ctx context.Context, eventID uuid.UUID) (*storage.Event, error) {
	event, err := a.storage.GetEvent(ctx, eventID)
	if err != nil {
		return nil, fmt.Errorf("get event from storage: %w", err)
	}
	return event, nil
}

func (a *App) GetEventsOfDay(ctx context.Context, since time.Time) ([]storage.Event, error) {
	return a.storage.GetEventsForPeriod(ctx, since, since.AddDate(0, 0, 1))
}

func (a *App) GetEventsOfWeek(ctx context.Context, since time.Time) ([]storage.Event, error) {
	return a.storage.GetEventsForPeriod(ctx, since, since.AddDate(0, 0, 7))
}

func (a *App) GetEventsOfMonth(ctx context.Context, since time.Time) ([]storage.Event, error) {
	return a.storage.GetEventsForPeriod(ctx, since, since.AddDate(0, 1, 0))
}
