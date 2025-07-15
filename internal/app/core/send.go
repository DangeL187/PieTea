package core

import (
	"fmt"
	"strings"

	"github.com/DangeL187/erax/pkg/erax"

	"PieTea/internal/app/request"
	"PieTea/internal/app/response"
	"PieTea/internal/infra/logger"
	"PieTea/internal/shared/exec"
	"PieTea/internal/shared/formatter"
)

// Send reads an HTTP request definition from a YAML file, executes the request
// using the `httpie` CLI tool, and returns the response headers and formatted body.
//
// The function performs the following steps:
//  1. Parses the YAML file at `filepath` into a Request object.
//  2. Converts the Request into CLI arguments for the `http` command,
//     adding flags to ignore stdin and print headers and body.
//  3. Executes the HTTP request via the external command.
//  4. Parses the raw HTTP response into separate headers and body strings.
//  5. Formats the response body with syntax highlighting.
//
// Parameters:
//   - filepath: path to the YAML file containing the HTTP request definition.
//
// Returns:
//   - headers: the raw HTTP response headers as a string.
//   - formattedBody: the pretty-printed, colorized HTTP response body.
//   - error: if any step fails (parsing, execution, formatting), an error is returned.
func Send(filepath string, showCmd bool) (string, string, string, erax.Error) {
	req, err := request.FromYAML(filepath)
	if err != nil {
		return "", "", "", erax.New(err, "Failed to create request from YAML").
			WithMeta("user_message", fmt.Sprintf("Failed to parse YAML:\n\n%v", err.Error()))
	}

	reqArgs := request.ToArgs(req)
	args := append([]string{"--ignore-stdin", "--print=hb"}, reqArgs...)

	var command string
	if showCmd {
		command = "http " + strings.Join(args, " ")
	}

	output, err := exec.Command("http", args...)
	if err != nil {
		return "", "", "", erax.New(err, "Failed to execute command")
	}

	headers, body := response.Parse(output)

	formattedBody, err := formatter.FormatJson(body)
	if err != nil {
		wrapped := erax.New(err, "Failed to format body").
			WithMeta("user_message", "Failed to format response body")
		logger.Logger.Warn("\n" + erax.Trace(wrapped))
	}

	return headers, formattedBody, command, nil
}
