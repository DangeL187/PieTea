package request

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Request struct {
	Method      string            `yaml:"Method"`
	URL         string            `yaml:"URL"`
	Headers     map[string]string `yaml:"Headers"`
	Body        map[string]string `yaml:"Body"`
	QueryParams map[string]string `yaml:"QueryParams"`
}

// FromYAML reads a YAML file from the given filepath and unmarshals it into a Request struct.
//
// For more details on how to write YAML, see the "usage" section in the README:
// https://github.com/DangeL187/PieTea?tab=readme-ov-file#usage
//
// Parameters:
//   - filepath: the path to the YAML file.
//
// Returns:
//   - A Request struct populated with the parsed data.
//   - An error if the file cannot be read or the YAML is invalid.
func FromYAML(filepath string) (Request, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return Request{}, fmt.Errorf("error reading file: %v", err)
	}

	var req Request
	err = yaml.Unmarshal(file, &req)
	if err != nil {
		return Request{}, fmt.Errorf("error unmarshaling YAML: %v", err)
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
