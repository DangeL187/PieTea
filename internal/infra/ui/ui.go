package ui

import (
	"fmt"
	"golang.org/x/term"
	"io"
	"os"

	"github.com/DangeL187/erax/pkg/erax"
	"github.com/charmbracelet/lipgloss"

	"PieTea/internal/app/response"
	"PieTea/internal/shared/config"
)

// Initializes output style.
//
// It retrieves the terminal width to dynamically set the output width, ensuring
// the content fits neatly within the styled border.
//
// Returns:
//   - An error if it fails to determine the terminal size.
func initOutput() erax.Error {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return erax.New(err, "Failed to get terminal size")
	}

	outputStyle = outputStyle.Width(width - outputStyle.GetHorizontalBorderSize())

	return nil
}

// Render outputs the given HTTP response headers and body to the terminal,
// formatting them within a rounded border that adapts to the current terminal width.
//
// Parameters:
//   - headers: the HTTP response headers as a string.
//   - body: the formatted HTTP response body as a string.
func Render(cfg config.Config, resp response.Response) {
	commandContent := ""
	if resp.Command != "" {
		commandContent = "Command: " + resp.Command
	}

	content := resp.Headers + "\n\n" + resp.Body

	err := initOutput()
	if err != nil || cfg.IsPlain {
		if commandContent != "" {
			fmt.Println(commandContent)
		}

		fmt.Println(content)

		return
	}

	if commandContent != "" {
		fmt.Println(outputStyle.Render(commandContent))
	}

	fmt.Println(outputStyle.Render(content))
}

// RenderError outputs a formatted error message to the given writer,
// formatting it within a rounded border that adapts to the current terminal width.
//
// Parameters:
//   - w: the destination writer (e.g., os.Stderr).
//   - format: an fmt.Sprintf-compatible format string.
//   - a: variadic arguments for formatting.
//
// Returns:
//   - An error if terminal size retrieval or writing to the output fails.
//   - Otherwise, returns nil after writing the styled error to the writer.
func RenderError(w io.Writer, format string, a ...any) erax.Error {
	err := initOutput()
	if err != nil {
		return erax.New(err, "Failed to initialize output")
	}

	content := fmt.Sprintf(format, a...)

	_, err2 := fmt.Fprint(w, outputStyle.Render(errorStyle.Render(content)))
	if err2 != nil {
		return erax.New(err2, "Failed to print rendered error")
	}

	return nil
}

var (
	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f38ba8"))
	outputStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder())
)
