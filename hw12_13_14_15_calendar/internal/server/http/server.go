package internalhttp

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"mime"
	"net/http"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/app"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/logger"
	pb "github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/pkg/calendarpb"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/third_party"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	cfg    *config.Calendar
	app    *app.App
	logger *logger.Logger
	srv    *http.Server
}

func NewServer(cfg *config.Calendar, app *app.App, logger *logger.Logger) *Server {
	return &Server{
		cfg:    cfg,
		app:    app,
		logger: logger,
	}
}

func (s *Server) Start(ctx context.Context) error {
	conn, err := grpc.DialContext(
		ctx, s.cfg.GRPC.Addr(),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("dial gRPC server: %w", err)
	}

	gwmux := runtime.NewServeMux()
	if err := pb.RegisterCalendarApiHandler(ctx, gwmux, conn); err != nil {
		return fmt.Errorf("reagister gateway: %w", err)
	}

	openAPI := openAPIHandler()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			gwmux.ServeHTTP(w, r)
			return
		}
		openAPI.ServeHTTP(w, r)
	})

	s.srv = &http.Server{
		Addr:        s.cfg.HTTP.Addr(),
		ReadTimeout: 5 * time.Second,
		Handler:     loggingMiddleware(mux),
	}

	s.logger.Info(fmt.Sprintf("http server is starting on %s", s.srv.Addr))
	if err := s.srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("serve http: %w", err)
		}
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if s.srv != nil {
		return s.srv.Shutdown(ctx)
	}
	return nil
}

func openAPIHandler() http.Handler {
	if err := mime.AddExtensionType(".svg", "image/svg+xml"); err != nil {
		logger.Error(fmt.Sprintf("failed to associated mime type: %v", err))
	}
	subFS, err := fs.Sub(third_party.OpenAPI, "OpenAPI")
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create sub filesystem: %v", err))
	}
	return http.FileServer(http.FS(subFS))
}
