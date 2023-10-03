package logging

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger = zerolog.Logger

func New() zerolog.Logger {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.999"
	return zerolog.New(os.Stderr).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Caller().
		Logger()
}

func NewConsoleLogger() zerolog.Logger {
	return zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05.999",
	}).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Caller().
		Logger()
}
