package cli

import (
	"flag"

	"github.com/DangeL187/erax/pkg/erax"
)

// ParseArgs validates and parses the command-line arguments.
//
// It expects exactly one argument: the path to a file.
// If the number of arguments is incorrect, it returns an error
// indicating the correct usage.
//
// Returns:
//   - A string containing the provided filepath.
//   - An error if the number of arguments is not equal to 1.
func ParseArgs() (string, bool, erax.Error) {
	showCmd := flag.Bool("show-cmd", false, "Показать сгенерированную команду перед выполнением")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		return "", false, erax.NewFromString("Arguments count mismatch", "").
			WithMeta("user_message", "Usage: ptea <filepath>")
	}
	filepath := args[0]
	return filepath, *showCmd, nil
}
