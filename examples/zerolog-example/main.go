package main

import "zerolog-example/logging"

func main() {
	logger := logging.New()
	logger.Info().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")
	logger.Error().
		Str("Name", "Tom").
		Send()

	ConsoleLogger := logging.NewConsoleLogger()
	ConsoleLogger.Info().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")
	ConsoleLogger.Error().
		Str("Name", "Tom").
		Send()
}
