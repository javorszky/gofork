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

const subs = 3

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	l := log.With().Str("who", "main").Logger()

	wg := sync.WaitGroup{}
	channels := make([]chan struct{}, subs)

	l.Info().Msgf("spinning up %d subs that will loop", subs)

	for i := 0; i < subs; i++ {
		go func(i int) {
			name := uuid.New().String()
			wg.Add(1)
			c := make(chan struct{}, 1)
			channels[i] = c
			app.New(name, l, c).Start()

			wg.Done()
			l.Info().Msgf("wg done called for %s", name)
		}(i)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-shutdown:
		log.Info().Msgf("main signal received: %s", sig.String())
		log.Info().Msg("prep main shutdown")
		defer log.Info().Msg("shutdown complete")
		log.Info().Msg("sending control signals to stopchans")
		for _, c := range channels {
			c <- struct{}{}
		}
		break
	}

	l.Info().Msg("waiting until wg.wait is done")

	wg.Wait()
}
