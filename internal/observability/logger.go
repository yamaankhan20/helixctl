// Package observability provides structured logging foundation.
package observability

import (
	"log/slog"
	"os"
)

// NewLogger creates a new structured logger for a component.
func NewLogger(component string) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("component", component)
}
