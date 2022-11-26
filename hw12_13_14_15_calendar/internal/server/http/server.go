package internalhttp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/app"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/logger"
)

type Server struct {
	logger *logger.Logger
	app    *app.App
	srv    *http.Server
}

func NewServer(logger *logger.Logger, app *app.App, addr string) *Server {
	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(hello))

	srv := &http.Server{
		Addr:        addr,
		ReadTimeout: 5 * time.Second,
		Handler:     loggingMiddleware(mux),
	}

	return &Server{
		logger: logger,
		app:    app,
		srv:    srv,
	}
}

func (s *Server) Start(ctx context.Context) error {
	if err := s.srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func hello(w http.ResponseWriter, _ *http.Request) {
	if _, err := io.WriteString(w, "Hello World!"); err != nil {
		logger.Global.Error(fmt.Sprintf("failed to write response: %v", err))
	}
}
