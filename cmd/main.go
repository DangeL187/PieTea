package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/DangeL187/erax/pkg/erax"
	"github.com/charmbracelet/log"

	"PieTea/cli"
	"PieTea/core"
	"PieTea/ui"
)

var logger = log.NewWithOptions(os.Stdout, log.Options{
	ReportTimestamp: true,
	TimeFormat:      time.TimeOnly,
})

func main() {
	logger.SetOutput(io.Discard) // TODO: enable with --debug flag (--log-file <filepath>)

	filepath, err := cli.ParseArgs()
	if err != nil {
		handleArgError(err)
	}

	headers, body, err := core.Send(filepath)
	if err != nil {
		handleSendError(err)
	}

	if err := ui.Render(headers, body); err != nil {
		handleRenderError(err)
	}
}

// --- Error Handlers ---

func handleArgError(err erax.Error) {
	logger.Error("\n" + erax.Trace(err))

	if msg, err := err.Meta("user_message"); err == nil {
		fmt.Println(msg)
	} else {
		fmt.Println("Invalid arguments")
	}

	os.Exit(0)
}

func handleSendError(err erax.Error) {
	wrapped := erax.New(err, "Failed to send request")
	logger.Error("\n" + erax.Trace(wrapped))

	msg, err := wrapped.Meta("user_message")
	if err != nil {
		msg = "Something went wrong"
	}

	if err := ui.RenderError(os.Stderr, "%v", msg); err != nil {
		errRender := erax.New(err, "Failed to render error")
		logger.Error("\n" + erax.Trace(errRender))

		_, _ = fmt.Fprintf(os.Stderr, "%v\n", msg)
	}

	os.Exit(0)
}

func handleRenderError(err erax.Error) {
	wrapped := erax.New(err, "Failed to render")
	logger.Error("\n" + erax.Trace(wrapped))

	// TODO: Fallback to plain output (--no-borders, --no-colors), depends on https://github.com/DangeL187/PieTea/issues/3

	_, _ = fmt.Fprintf(os.Stderr, "%s\n", wrapped.Msg())
	os.Exit(0)
}
