package config

// Config holds configuration options for the application.
type Config struct {
	Filepath          string
	IgnoreMissingVars bool
	IsDebug           bool
	LogFile           string
	ShowCmd           bool
}
