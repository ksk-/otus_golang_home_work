package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/logger"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/rmq"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/schedule"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/storage"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/scheduler.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	cfg, err := config.NewSchedulerConfig(configFile)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to configure app: %v", err))
		os.Exit(1)
	}

	l := logger.New(&cfg.Logger).WithGlobal()
	s, err := storage.NewStorage(&cfg.Storage, l)
	if err != nil {
		l.Error(fmt.Sprintf("faield to create storage: %v", err))
		os.Exit(1)
	}
	defer func() {
		if err := s.Close(); err != nil {
			l.Error(fmt.Sprintf("faield to close storage: %v", err))
		}
	}()

	notifier, err := rmq.NewNotifier(cfg.RMQ.URI(), cfg.RMQ.Queue)
	if err != nil {
		l.Error(fmt.Sprintf("failed to creqte RMQ notifier: %v", err))
		os.Exit(1) //nolint:gocritic
	}
	defer func() {
		if err := notifier.Close(); err != nil {
			l.Error(fmt.Sprintf("faield to close RMQ notifier: %v", err))
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	l.Info("scheduler is running...")
	scheduler := schedule.NewScheduler(s, notifier)

	from := time.Now()
	to := from.Add(cfg.Tick)
	scheduler.CheckCalendar(ctx, from, to)

	ticker := time.NewTicker(cfg.Tick)
	defer ticker.Stop()

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case <-ticker.C:
			from, to = from.Add(cfg.Tick), to.Add(cfg.Tick)
			scheduler.CheckCalendar(ctx, from, to)
		}
	}
	<-ctx.Done()
	scheduler.Wait()
	l.Info("scheduler is stopping...")
}
