package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/logger"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/notify"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/rmq"
)

const sentNotificationsTopic = "event_notifications_sent"

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/sender.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	cfg, err := config.NewSenderConfig(configFile)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to configure app: %v", err))
		os.Exit(1)
	}

	logger.New(&cfg.Logger).WithGlobal()

	queue, err := rmq.NewQueue(cfg.RMQ.URI(), cfg.RMQ.Queue)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create RMQ queue: %v", err))
		os.Exit(1)
	}
	defer func() {
		if err := queue.Close(); err != nil {
			logger.Error(fmt.Sprintf("failed to close RMQ queue: %v", err))
		}
	}()

	notifier, err := rmq.NewNotifier(cfg.RMQ.URI(), sentNotificationsTopic)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create RMQ notifier: %v", err))
		os.Exit(1) //nolint:gocritic
	}
	defer func() {
		if err := notifier.Close(); err != nil {
			logger.Error(fmt.Sprintf("failed to close RMQ notifier: %v", err))
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	ch, err := queue.ConsumeChannel(ctx, "calendar_sender")
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create consume channel: %v", err))
		os.Exit(1)
	}

	logger.Info("sender is running...")
	for msg := range ch {
		var notification notify.Notification
		if err := json.Unmarshal(msg, &notification); err != nil {
			logger.Error(fmt.Sprintf("failed to unmarshal notification: %v", err))
			continue
		}
		if err := notifier.Notify(ctx, &notification); err != nil {
			logger.Error(fmt.Sprintf("failed to send notificaion %v: %v", notification.ID, err))
		}
	}
	<-ctx.Done()
	logger.Info("sender is stopping...")
}
