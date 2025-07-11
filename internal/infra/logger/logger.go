package logger

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var Logger = log.NewWithOptions(os.Stdout, log.Options{
	ReportCaller:    true,
	ReportTimestamp: true,
	TimeFormat:      time.TimeOnly,
})
