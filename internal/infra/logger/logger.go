package logger

import (
	"io"
	"os"
	"time"

	"github.com/DangeL187/erax/pkg/erax"
	"github.com/charmbracelet/log"

	"PieTea/internal/shared/config"
)

// Logger is a global logger with configurable output and formatting options.
var Logger = log.NewWithOptions(os.Stdout, log.Options{
	ReportCaller:    true,
	ReportTimestamp: true,
	TimeFormat:      time.TimeOnly,
})

// Init configures the global Logger based on the given config.
//
// If debug mode is disabled, logging is suppressed.
// If a log file path is provided, logs are written to that file.
//
// Returns an error if the log file cannot be opened.
func Init(cfg config.Config) erax.Error {
	if !cfg.IsDebug {
		Logger.SetOutput(io.Discard)
	}

	if cfg.LogFile != "" {
		logFile, err := os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return erax.New(err, "Failed to open log file")
		}

		Logger.SetOutput(logFile)
	}

	return nil
}
