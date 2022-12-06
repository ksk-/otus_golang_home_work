package schedule

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/logger"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/notify"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/storage"
)

type Scheduler struct {
	storage  storage.Storage
	notifier notify.Notifier
	wg       sync.WaitGroup
}

func (s *Scheduler) CheckCalendar(ctx context.Context, from, to time.Time) {
	deleted, err := s.storage.DeletePastEvents(ctx, from.AddDate(-1, 0, 0))
	if err != nil {
		logger.Error(fmt.Sprintf("failed to delete past events: %v", err))
		return
	}
	if deleted > 0 {
		logger.Info(fmt.Sprintf("deleted %d past events", deleted))
	}

	events, err := s.storage.GetEventsToNotify(ctx, from, to)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get events: %v", err))
		return
	}
	s.scheduleNotifications(ctx, events)
}

func (s *Scheduler) Wait() {
	s.wg.Wait()
}

func (s *Scheduler) scheduleNotifications(ctx context.Context, events []storage.Event) {
	s.wg.Add(len(events))
	for _, event := range events {
		go func(event storage.Event) {
			defer s.wg.Done()
			select {
			case <-ctx.Done():
				logger.Warn(fmt.Sprintf("event %s was scheduled, but no notified", event.ID))
				return
			case <-time.After(time.Until(event.NotificationTime)):
				s.notifyAboutEvent(ctx, &event)
				return
			}
		}(event)
	}
}

func (s *Scheduler) notifyAboutEvent(ctx context.Context, event *storage.Event) {
	notification := notify.Notification{
		ID:        uuid.New(),
		EventID:   event.ID,
		Title:     event.Title,
		EventTime: event.BeginTime,
		UserID:    event.UserID,
	}
	if err := s.notifier.Notify(ctx, &notification); err != nil {
		logger.Error(fmt.Sprintf("failed to notify about event %s", event.ID))
	} else {
		logger.Debug(fmt.Sprintf("[notification sent]: ID: '%v' EventID: '%v'", notification.ID, notification.EventID))
	}
}

func NewScheduler(storage storage.Storage, notifier notify.Notifier) *Scheduler {
	return &Scheduler{
		storage:  storage,
		notifier: notifier,
	}
}
