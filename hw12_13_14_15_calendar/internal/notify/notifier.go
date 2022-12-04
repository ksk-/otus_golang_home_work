package notify

import (
	"context"
	"fmt"

	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/logger"
)

type Notifier interface {
	Notify(ctx context.Context, notification *Notification) error
}

func NewLogNotifier(logger *logger.Logger) Notifier {
	return &logNotifier{logger: logger}
}

type logNotifier struct {
	logger *logger.Logger
}

func (n *logNotifier) Notify(_ context.Context, notification *Notification) error {
	n.logger.Info(fmt.Sprintf("[UPCOMING EVENT NOTIFICATION]: %v", *notification))
	return nil
}
