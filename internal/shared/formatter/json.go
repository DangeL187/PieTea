package formatter

import (
	"encoding/json"

	"github.com/DangeL187/erax/pkg/erax"
	"github.com/TylerBrock/colorjson"
)

// FormatJson attempts to pretty-print a raw JSON HTTP response body
// using syntax highlighting and indentation.
//
// It uses the github.com/TylerBrock/colorjson package to apply formatting.
// If the input is not valid JSON, an error is returned.
//
// Parameters:
//   - body: a string containing raw JSON (usually from an HTTP response).
//
// Returns:
//   - a syntax-highlighted, indented JSON string suitable for terminal output.
//   - an error if the input is not valid JSON or cannot be formatted.
func FormatJson(rawJson string) (string, erax.Error) {
	formatter := colorjson.NewFormatter()
	formatter.Indent = 2

	var obj interface{}
	err := json.Unmarshal([]byte(rawJson), &obj)
	if err != nil {
		return "", erax.New(err, "Failed to unmarshal JSON")
	}

	formattedBody, err := formatter.Marshal(obj)
	if err != nil {
		return "", erax.New(err, "Failed to marshal JSON")
	}

	return string(formattedBody), nil
}
