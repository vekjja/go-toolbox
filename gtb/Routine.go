package gtb

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

type Routine struct {
}

// NewRoutine creates a new Routine instance.
func NewRoutine(_default, _defer func()) context.Context {
	// Create a context that is canceled when the command is interrupted or completed
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle interrupt signals for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		cancel()
	}()

	go func() {
		defer _defer()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_default()
			}
		}
	}()
	// Block until the context is canceled
	// <-ctx.Done()
	return ctx
}
