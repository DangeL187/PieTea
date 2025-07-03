package core

import (
	"fmt"

	"github.com/DangeL187/erax/pkg/erax"

	"PieTea/core/request"
	"PieTea/util"
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
func Send(filepath string) (string, string, erax.Error) {
	req, err := request.FromYAML(filepath)
	if err != nil {
		return "", "", erax.New(err, "Failed to create request from YAML").
			WithMeta("user_message", fmt.Sprintf("Failed to parse YAML:\n\n%v", err.Error()))
	}

	reqArgs := request.ToArgs(req)
	args := append([]string{"--ignore-stdin", "--print=hb"}, reqArgs...)

	response, err := util.ExecCommand("http", args...)
	if err != nil {
		return "", "", erax.New(err, "Failed to execute command")
	}

	headers, body := util.ParseHttpResponse(response)

	formattedBody, err := util.FormatHttpBody(body)
	if err != nil {
		return "", "", erax.New(err, "Failed to format body").
			WithMeta("user_message", "Failed to format response body")
	}

	return headers, formattedBody, nil
}
