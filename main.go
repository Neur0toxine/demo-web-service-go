package main

import (
	"os"

	"github.com/Neur0toxine/demo-web-service-go/cmd"
	"github.com/jessevdk/go-flags"
)

func main() {
	if _, err := cmd.Parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		os.Exit(-1)
	}
}
