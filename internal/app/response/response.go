package response

import "strings"

// Parse splits a raw HTTP response string into headers and body parts.
//
// It first attempts to split the response by the standard HTTP header-body separator "\r\n\r\n".
// If that fails, it falls back to splitting by a double newline "\n\n".
//
// Parameters:
//   - response: the full raw HTTP response as a single string.
//
// Returns:
//   - headers: the HTTP response headers as a string.
//   - body: the HTTP response body as a string (empty if not present)
func Parse(response string) (string, string) {
	parts := strings.SplitN(response, "\r\n\r\n", 2)
	if len(parts) < 2 {
		parts = strings.SplitN(response, "\n\n", 2)
	}

	headers := parts[0]
	body := ""
	if len(parts) == 2 {
		body = parts[1]
	}

	return headers, body
}
