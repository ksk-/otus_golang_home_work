package internalgrpc

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

const timeFormat = "02/Jan/2006:15:04:05 -0700"

func serverUnaryInterceptor(
	ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (any, error) {
	reqTime := time.Now()
	m, err := handler(ctx, req)

	host := clientHost(ctx)
	msg := fmt.Sprintf("%s [%s] %s %s ", host, reqTime.Format(timeFormat), info.FullMethod, time.Since(reqTime))
	logger.Debug(msg)

	return m, err
}

func clientHost(ctx context.Context) string {
	if p, ok := peer.FromContext(ctx); ok {
		if host, _, err := net.SplitHostPort(p.Addr.String()); err == nil {
			return host
		}
	}
	return "<unknown host>"
}
