package request

import (
	"gopkg.in/yaml.v3"
	"os"
	"strings"

	"github.com/DangeL187/erax/pkg/erax"

	"PieTea/internal/shared/config"
)

// Request represents an HTTP request definition.
//
// Fields:
//   - Method: HTTP method (e.g., GET, POST).
//   - URL: request URL.
//   - Headers: HTTP headers as key-value pairs.
//   - Body: request body as key-value pairs.
//   - QueryParams: URL query parameters as key-value pairs.
type Request struct {
	Method      string            `yaml:"Method"`
	URL         string            `yaml:"URL"`
	Headers     map[string]string `yaml:"Headers"`
	Body        map[string]string `yaml:"Body"`
	QueryParams map[string]string `yaml:"QueryParams"`
}

// FromYAML reads a YAML file from the given filepath, substitutes ${VAR} placeholders
// with environment variable values, and unmarshalls the result into a Request.
//
// Placeholders in the form ${VAR} are replaced by the value of the corresponding
// environment variable. If the variable is not set, the placeholder remains unchanged.
//
// For usage details, see the README: https://github.com/DangeL187/PieTea?tab=readme-ov-file#usage
//
// Parameters:
//   - cfg: configuration containing the YAML filepath.
//
// Returns:
//   - Request populated from the processed YAML.
//   - erax.Error if reading or parsing fails.
func FromYAML(cfg config.Config) (Request, erax.Error) {
	content, err := os.ReadFile(cfg.Filepath)
	if err != nil {
		return Request{}, erax.New(err, "Error reading file")
	}

	var missingVars []string

	processed := os.Expand(string(content), func(key string) string {
		value := os.Getenv(key)
		if value == "" {
			if !cfg.IgnoreMissingVars {
				missingVars = append(missingVars, key)
			}
			return "${" + key + "}"
		}

		return value
	})

	if len(missingVars) > 0 {
		return Request{}, erax.NewFromString(
			"Environment variables not set: "+strings.Join(missingVars, ", "),
			"",
		)
	}

	var req Request
	err = yaml.Unmarshal([]byte(processed), &req)
	if err != nil {
		return Request{}, erax.New(err, "Error unmarshalling YAML")
	}

	return req, nil
}

// ToArgs converts a Request into a slice of strings formatted as command-line arguments.
//
// This function is intended to prepare arguments for an HTTP client CLI,
// such as httpie. It formats headers as "Key:Value", body fields as "Key=Value",
// and query parameters as "Key::Value". These are appended to the method and URL.
//
// Parameters:
//   - req: the Request to be converted.
//
// Returns:
//   - A slice of strings representing command-line arguments.
func ToArgs(req Request) []string {
	args := []string{req.Method, req.URL}

	for key, value := range req.Headers {
		args = append(args, key+":"+value)
	}

	for key, value := range req.Body {
		args = append(args, key+"="+value)
	}

	for key, value := range req.QueryParams {
		args = append(args, key+"::"+value)
	}

	return args
}
