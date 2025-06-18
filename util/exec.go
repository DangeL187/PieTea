package util

import "os/exec"

// ExecCommand runs an external command specified by cmdName and its arguments args,
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
//   - If the command fails (non-zero exit status), the output is still returned along with the error.
func ExecCommand(cmdName string, args ...string) (string, error) {
	cmd := exec.Command(cmdName, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}
	return string(output), nil
}
