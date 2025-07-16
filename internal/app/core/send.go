package core

import (
	"PieTea/internal/shared/config"
	"fmt"
	"strings"

	"github.com/DangeL187/erax/pkg/erax"

	"PieTea/internal/app/request"
	"PieTea/internal/app/response"
	"PieTea/internal/infra/logger"
	"PieTea/internal/shared/exec"
	"PieTea/internal/shared/formatter"
)

// Send loads an HTTP request from a YAML file, runs it via `httpie` CLI, and returns the response.
//
// It parses the YAML into a request, builds CLI args, runs the command, and formats the response.
//
// Returns:
//   - response.Response with headers, formatted body, and optional executed command string.
//   - erax.Error if parsing, execution, or formatting fails.
func Send(cfg config.Config) (response.Response, erax.Error) {
	req, err := request.FromYAML(cfg)
	if err != nil {
		return response.Response{}, erax.New(err, "Failed to create request from YAML").
			WithMeta("user_message", fmt.Sprintf("Failed to parse YAML:\n\n%v", err.Error()))
	}

	reqArgs := request.ToArgs(req)
	args := append([]string{"--ignore-stdin", "--print=hb"}, reqArgs...)

	var command string
	if cfg.ShowCmd {
		command = "http " + strings.Join(args, " ")
	}

	output, isErr := exec.Command("http", args...)
	if isErr {
		return response.Response{}, erax.New(err, "Failed to execute command").
			WithMeta("user_message", output)
	}

	headers, body := response.Parse(output)

	formattedBody, err := formatter.FormatJSON(body)
	if err != nil {
		wrapped := erax.New(err, "Failed to format body").
			WithMeta("user_message", "Failed to format response body")
		logger.Logger.Warn("\n" + erax.Trace(wrapped))
	}

	resp := response.Response{
		Body:    formattedBody,
		Command: command,
		Headers: headers,
	}

	return resp, nil
}
