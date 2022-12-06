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
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/app"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/logger"
	internalgrpc "github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/server/grpc"
	internalhttp "github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/server/http"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/storage"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/calendar.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	cfg, err := config.NewCalendarConfig(configFile)
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

	application := app.New(s, l)
	grpcSrv := internalgrpc.NewServer(cfg, application, l)
	httpSrv := internalhttp.NewServer(cfg, application, l)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	l.Info("calendar is running...")
	go func() {
		if err := grpcSrv.Start(); err != nil {
			l.Error(fmt.Sprintf("failed to start gRPC server: %v", err))
			cancel()
			os.Exit(1)
		}
	}()
	go func() {
		if err := httpSrv.Start(ctx); err != nil {
			l.Error(fmt.Sprintf("failed to start http server: %v", err))
			cancel()
			os.Exit(1)
		}
	}()
	<-ctx.Done()

	l.Info("calendar is stopping...")
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := httpSrv.Stop(ctx); err != nil {
		l.Error("failed to stop http server: " + err.Error())
	}
	grpcSrv.Stop()
}
