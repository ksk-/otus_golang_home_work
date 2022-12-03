package schedule

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/logger"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/notify"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/storage"
)

type Scheduler struct {
	cfg      *config.Scheduler
	storage  storage.Storage
	notifier notify.Notifier
	wg       sync.WaitGroup
}

func (s *Scheduler) LoadSchedule(ctx context.Context) {
	now := time.Now()
	deleted, err := s.storage.DeletePastEvents(ctx, now.AddDate(-1, 0, 0))
	if err != nil {
		logger.Error(fmt.Sprintf("failed to delete past events: %v", err))
		return
	}
	if deleted > 0 {
		logger.Info(fmt.Sprintf("deleted %d past events", deleted))
	}

	events, err := s.storage.GetEventsToNotify(ctx, now, now.Add(s.cfg.Tick))
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get events: %v", err))
		return
	}
	s.scheduleNotification(ctx, events)
}

func (s *Scheduler) Wait() {
	s.wg.Wait()
}

func (s *Scheduler) scheduleNotification(ctx context.Context, events []storage.Event) {
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
	}
}

func NewScheduler(cfg *config.Scheduler, storage storage.Storage, notifier notify.Notifier) *Scheduler {
	return &Scheduler{
		cfg:      cfg,
		storage:  storage,
		notifier: notifier,
	}
}
