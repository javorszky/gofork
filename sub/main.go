package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/google/uuid"
	"github.com/javorszky/sub/app"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	l := log.With().Str("who", "main").Logger()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	stopChan := make(chan struct{}, 1)

	wg := sync.WaitGroup{}

	go func() {
		wg.Add(1)
		name := uuid.New().String()
		l.Info().Msgf("starting app with name %s", name)

		// this will wait until Start returns. Start only returns when the stopchan gets an input.
		app.New(name, l, stopChan).Start()

		l.Info().Msgf("app with name %s stopped", name)
		wg.Done()
	}()

	select {
	case sig := <-shutdown:
		l.Info().Msgf("signal received: %s", sig.String())
		l.Info().Msg("prep shutdown")
		defer l.Info().Msg("shutdown complete")
		l.Info().Msg("sending control signal to stopchan")
		stopChan <- struct{}{}
	}

	wg.Wait()
}
