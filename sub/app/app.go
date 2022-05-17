package app

import (
	"time"

	"github.com/rs/zerolog"
)

type App struct {
	stop   <-chan struct{}
	name   string
	logger zerolog.Logger
}

func New(name string, logger zerolog.Logger, stop <-chan struct{}) App {
	l := logger.With().Str("where", "sub").Str("name", name).Logger()
	return App{
		name:   name,
		stop:   stop,
		logger: l,
	}
}

func (a App) Start() {
loop:
	for {
		select {
		case <-time.After(time.Second):
			a.logger.Info().Msgf("another loop bites the dust")
		case <-a.stop:
			a.logger.Info().Msg("stop control signal received, breaking out of the for loop, terminating go func")
			break loop
		}
	}

	a.logger.Info().Msg("releasing control")
}
