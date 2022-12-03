package internalhttp

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/logger"
)

const timeFormat = "02/Jan/2006:15:04:05 -0700"

type wrappedResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *wrappedResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer := &wrappedResponseWriter{ResponseWriter: w, status: http.StatusOK}

		reqTime := time.Now()
		next.ServeHTTP(writer, r)

		host := clientHost(r)
		status := http.StatusText(writer.status)
		msg := fmt.Sprintf(
			"%s [%s] %s %s %s %s %s %s",
			host, reqTime.Format(timeFormat), r.Method, r.URL, r.Proto, status, time.Since(reqTime), r.UserAgent(),
		)

		logger.Debug(msg)
	})
}

func clientHost(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "<unknown host>"
	}
	return host
}
