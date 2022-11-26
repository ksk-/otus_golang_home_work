package logger

import (
	"encoding/json"
	"io"
	"os"
	"testing"
	"time"

	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
		output := runAndReadOutput(t, func() {
			l := New(&config.Logger{Level: config.InfoLogLevel, Pretty: false})
			l.Info("record msg")
		})
		record := make(map[string]interface{})
		require.NoError(t, json.Unmarshal(output, &record))

		timestamp, err := time.Parse(time.RFC3339, record["time"].(string))
		require.NoError(t, err)
		require.WithinDuration(t, timestamp, time.Now(), time.Second)
		require.Equal(t, "info", record["level"])
		require.Equal(t, "record msg", record["message"])
	})
	t.Run("less log level", func(t *testing.T) {
		output := runAndReadOutput(t, func() {
			l := New(&config.Logger{Level: config.WarnLogLevel, Pretty: false})
			l.Debug("debug record msg")
			l.Info("info record msg")
		})
		require.Empty(t, output)

		output = runAndReadOutput(t, func() {
			l := New(&config.Logger{Level: config.DebugLogLevel, Pretty: false})
			l.Debug("debug record msg")
			l.Info("info record msg")
		})
		require.Contains(t, string(output), "debug record msg")
		require.Contains(t, string(output), "info record msg")
	})
	t.Run("pretty output", func(t *testing.T) {
		output := runAndReadOutput(t, func() {
			l := New(&config.Logger{Pretty: true})
			l.Info("record msg")
		})
		require.Contains(t, string(output), "INF", "record msg")
	})
}

func runAndReadOutput(t *testing.T, callable func()) []byte {
	t.Helper()

	original := os.Stdout
	defer func() {
		os.Stdout = original
	}()

	r, w, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = w
	callable()
	require.NoError(t, w.Close())

	output, err := io.ReadAll(r)
	require.NoError(t, err)
	require.NoError(t, r.Close())

	return output
}
