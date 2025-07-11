package exec

import (
	"os/exec"
	"strings"

	"github.com/DangeL187/erax/pkg/erax"
)

// Command runs an external command specified by cmdName and its arguments args,
// then returns the combined standard output and standard error as a string.
//
// Parameters:
//   - cmdName: the name or path of the executable to run.
//   - args: zero or more arguments to pass to the command.
//
// Returns:
//   - string: combined output (stdout and stderr) of the executed command.
//   - error: any error encountered while running the command.
//
// Notes:
//   - If the command fails (non-zero exit status), the output will be empty and message will be returned along with the error.
func Command(cmdName string, args ...string) (string, erax.Error) {
	cmd := exec.Command(cmdName, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", erax.New(err, "Failed to run CombinedOutput").
			WithMeta("user_message", strings.Trim(string(output), "\n"))
	}
	return string(output), nil
}
