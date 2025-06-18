package cli

import (
	"errors"
	"os"
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
func ParseArgs() (string, error) {
	if len(os.Args) != 2 {
		return "", errors.New("usage: ptea <filepath>")
	}

	return os.Args[1], nil
}
