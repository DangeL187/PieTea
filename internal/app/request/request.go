package request

import (
	"gopkg.in/yaml.v3"
	"os"

	"github.com/DangeL187/erax/pkg/erax"
)

type Request struct {
	Method      string            `yaml:"Method"`
	URL         string            `yaml:"URL"`
	Headers     map[string]string `yaml:"Headers"`
	Body        map[string]string `yaml:"Body"`
	QueryParams map[string]string `yaml:"QueryParams"`
}

// FromYAML reads a YAML file from the given filepath, substitutes ${VAR} placeholders
// with values from environment variables, and unmarshalls the result into a Request struct.
//
// The YAML file may include placeholders in the form ${VAR}, which will be replaced
// using the corresponding environment variables. If a variable is not set, the placeholder
// will remain unchanged.
//
// For more details on how to write YAML with variable placeholders, see the "usage" section
// in the README: https://github.com/DangeL187/PieTea?tab=readme-ov-file#usage
//
// Parameters:
//   - filepath: the path to the YAML file.
//
// Returns:
//   - A Request struct populated with the parsed and processed YAML data.
//   - An erax.Error if the file cannot be read or the YAML is invalid.
func FromYAML(filepath string) (Request, erax.Error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return Request{}, erax.New(err, "Error reading file")
	}

	processed := os.Expand(string(content), func(key string) string {
		value := os.Getenv(key)
		if value == "" {
			return "${" + key + "}"
		}

		return value
	})

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
