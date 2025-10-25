package main

import (
	"fmt"
	"os"

	"github.com/DangeL187/erax"

	"PieTea/internal/app/core"
	"PieTea/internal/cli"
	"PieTea/internal/infra/logger"
	"PieTea/internal/infra/ui"
)

func main() {
	cfg, err := cli.ParseArgs()
	if err != nil {
		os.Exit(1)
	}

	err = logger.Init(cfg)
	if err != nil {
		handleLoggerError(err)
	}

	resp, err := core.Send(cfg)
	if err != nil {
		handleSendError(err)
	}

	ui.Render(cfg, resp)
}

// --- Error Handlers ---

func handleLoggerError(err error) {
	wrapped := erax.New(err, "Failed to init logger")
	logger.Logger.Error("\n" + erax.Trace(wrapped))

	os.Exit(2)
}

func handleSendError(err error) {
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
