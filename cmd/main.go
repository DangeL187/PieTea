package main

import (
	"fmt"
	"os"

	"PieTea/cli"
	"PieTea/core"
	"PieTea/ui"
)

func main() {
	//logger := log.NewWithOptions(os.Stdout, log.Options{
	//	ReportTimestamp: true,
	//	TimeFormat:      time.TimeOnly,
	//})

	filepath, err := cli.ParseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	headers, body, err := core.Send(filepath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to send request: %v", err)
		os.Exit(1)
	}

	err = ui.Render(headers, body)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to render: %v", err)
		os.Exit(1)
	}
}
