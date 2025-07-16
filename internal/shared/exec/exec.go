package exec

import (
	"os/exec"
	"strings"
)

// Command runs an external command and returns its combined output.
//
// Parameters:
//   - cmdName: executable name or path.
//   - args: command arguments.
//
// Returns:
//   - combined stdout and stderr output.
//   - true if command failed (non-zero exit code or execution error), false otherwise.
func Command(cmdName string, args ...string) (string, bool) {
	cmd := exec.Command(cmdName, args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return strings.Trim(string(output), "\n"), true
	}

	return string(output), false
}
