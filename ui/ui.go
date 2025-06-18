package ui

import (
	"fmt"
	"golang.org/x/term"
	"os"

	"github.com/charmbracelet/lipgloss"
)

// Render outputs the given HTTP response headers and body to the terminal,
// formatting them within a rounded border that adapts to the current terminal width.
//
// It retrieves the terminal width to dynamically set the output width, ensuring
// the content fits neatly within the styled border.
//
// Parameters:
//   - headers: the HTTP response headers as a string.
//   - body: the formatted HTTP response body as a string.
//
// Returns:
//   - An error if it fails to determine the terminal size.
//   - Otherwise, returns nil after printing the styled output to stdout.
func Render(headers, body string) error {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return fmt.Errorf("failed to get terminal size: %v", err)
	}

	outputStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder())
	outputStyle = outputStyle.Width(width - outputStyle.GetHorizontalBorderSize())

	fmt.Println(outputStyle.Render(headers + "\n\n" + body))

	return nil
}
