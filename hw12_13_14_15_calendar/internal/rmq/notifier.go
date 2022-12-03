package rmq

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/notify"
)

type Notifier struct {
	queue *Queue
}

func (n *Notifier) Notify(ctx context.Context, notification *notify.Notification) error {
	msg, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("marshal to json: %w", err)
	}
	return n.queue.Push(ctx, msg)
}

func (n *Notifier) Close() error {
	return n.queue.Close()
}

func NewNotifier(cfg *config.RMQConfig) (*Notifier, error) {
	queue, err := NewQueue(cfg, "application/json")
	if err != nil {
		return nil, fmt.Errorf("create queue: %w", err)
	}
	return &Notifier{queue: queue}, nil
}
