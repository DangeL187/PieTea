package config

// Config holds configuration options for the application.
type Config struct {
	Filepath          string
	IgnoreMissingVars bool
	IsDebug           bool
	IsPlain           bool
	LogFile           string
	ShowCmd           bool
}
