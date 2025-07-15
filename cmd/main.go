package main

import (
	"fmt"
	"os"

	"github.com/DangeL187/erax/pkg/erax"

	"PieTea/internal/app/core"
	"PieTea/internal/cli"
	"PieTea/internal/infra/logger"
	"PieTea/internal/infra/ui"
)

func main() {
	//logger.Logger.SetOutput(io.Discard) // TODO: enable with --debug flag (--log-file <filepath>)

	filepath, showCmd, err := cli.ParseArgs()
	if err != nil {
		handleArgError(err)
	}

	headers, body, command, err := core.Send(filepath, showCmd)
	if err != nil {
		handleSendError(err)
	}

	ui.Render(headers, body, command)
}

// --- Error Handlers ---

func handleArgError(err erax.Error) {
	logger.Logger.Error("\n" + erax.Trace(err))

	if msg, err := err.Meta("user_message"); err == nil {
		fmt.Println(msg)
	} else {
		fmt.Println("Invalid arguments")
	}

	os.Exit(0)
}

func handleSendError(err erax.Error) {
	wrapped := erax.New(err, "Failed to send request")
	logger.Logger.Error("\n" + erax.Trace(wrapped))

	msg, err := wrapped.Meta("user_message")
	if err != nil {
		msg = "Something went wrong"
	}

	if err := ui.RenderError(os.Stderr, "%v", msg); err != nil {
		errRender := erax.New(err, "Failed to render error")
		logger.Logger.Error("\n" + erax.Trace(errRender))

		_, _ = fmt.Fprintf(os.Stderr, "%v\n", msg)
	}

	os.Exit(0)
}

func handleRenderError(err erax.Error) {
	wrapped := erax.New(err, "Failed to render")
	logger.Logger.Error("\n" + erax.Trace(wrapped))

	// TODO: Fallback to plain output (--no-borders, --no-colors), depends on https://github.com/DangeL187/PieTea/issues/3

	_, _ = fmt.Fprintf(os.Stderr, "%s\n", wrapped.Msg())
	os.Exit(0)
}
