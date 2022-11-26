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
	"github.com/jmoiron/sqlx"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/app"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/logger"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/server/http"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/storage"
)

const (
	memoryStorage = "memory"
	sqlStorage    = "sql"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	cfg, err := config.NewConfig(configFile)
	if err != nil {
		logger.Global.Error(fmt.Sprintf("failed to configure app: %v", err))
		os.Exit(1)
	}

	l := logger.New(&cfg.Logger)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	var s storage.Storage
	switch cfg.Storage.Type {
	case memoryStorage:
		s = storage.NewMemoryStorage()
	case sqlStorage:
		db, err := sqlx.Connect("pgx", cfg.Storage.DB.DSN())
		if err != nil {
			l.Error(fmt.Sprintf("failed to create DB connection: %v", err))
			os.Exit(1) //nolint:gocritic
		}
		defer func() {
			if err := db.Close(); err != nil {
				l.Error(fmt.Sprintf("faield to close DB connection: %v", err))
			}
		}()
	default:
		l.Error(fmt.Sprintf("unknown storage type: %s", cfg.Storage.Type))
		os.Exit(1)
	}

	server := internalhttp.NewServer(l, app.New(l, s), cfg.HTTP.Addr())

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			l.Error("failed to stop http server: " + err.Error())
		}
	}()

	l.Info("calendar is running...")
	if err := server.Start(ctx); err != nil {
		l.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1)
	}
	l.Info("calendar is stopping...")
}
