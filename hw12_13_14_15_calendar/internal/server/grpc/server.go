package internalgrpc

import (
	"fmt"
	"net"

	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/api"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/app"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/logger"
	pb "github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/pkg/calendarpb"
	"google.golang.org/grpc"
)

func NewServer(cfg *config.Calendar, app *app.App, logger *logger.Logger) *Server {
	srv := grpc.NewServer(grpc.UnaryInterceptor(serverUnaryInterceptor))
	pb.RegisterCalendarApiServer(srv, api.NewAPI(app))

	return &Server{
		addr:   cfg.GRPC.Addr(),
		srv:    srv,
		logger: logger,
	}
}

type Server struct {
	addr   string
	srv    *grpc.Server
	logger *logger.Logger
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("listen %s: %w", s.addr, err)
	}
	s.logger.Info(fmt.Sprintf("gRPC server is starting on %s", s.addr))
	if err = s.srv.Serve(listener); err != nil {
		return fmt.Errorf("gRPC serve: %w", err)
	}
	return nil
}

func (s *Server) Stop() {
	s.srv.GracefulStop()
}
