package cli

import (
	"flag"
	"fmt"

	"github.com/DangeL187/erax"

	"PieTea/internal/shared/config"
)

// printUsage displays CLI usage information and available flags.
func printUsage() {
	_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage: ptea [OPTIONS] [FILE]\nOptions:\n")
	flag.PrintDefaults()
}

// ParseArgs parses command-line flags and arguments.
//
// Returns:
//   - config.Config: populated configuration struct based on flags and arguments.
//   - error: returned if the required positional argument (filepath) is missing.
func ParseArgs() (config.Config, error) {
	ignoreMissingVars := flag.Bool("ignore-missing-vars", false, "Ignore missing or unset environment variables in the YAML file")
	isDebug := flag.Bool("debug", false, "Debug mode")
	logFile := flag.String("log-file", "", "Write logs to specified file instead of stdout")
	showCmd := flag.Bool("show-cmd", false, "Show generated command before executing")
	var isPlain bool
	flag.BoolVar(&isPlain, "p", false, "Only show plain style, no decorations (short)")
	flag.BoolVar(&isPlain, "plain", false, "Only show plain style, no decorations (long)")

	flag.Usage = printUsage

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		printUsage()
		return config.Config{}, erax.NewFromString("Arguments count mismatch", "")
	}

	cfg := config.Config{
		Filepath:          args[0],
		IgnoreMissingVars: *ignoreMissingVars,
		IsDebug:           *isDebug,
		IsPlain:           isPlain,
		LogFile:           *logFile,
		ShowCmd:           *showCmd,
	}

	return cfg, nil
}
