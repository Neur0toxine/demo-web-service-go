package cmd

import (
	"fmt"
	"time"

	"github.com/Neur0toxine/demo-web-service-go/internal/app"
	"github.com/brianvoe/gofakeit/v6"
)

func init() {
	_, err := Parser.AddCommand("run", "Run the application", "", &RunCommand{})
	if err != nil {
		panic(err)
	}
}

// RunCommand is used to run the server
type RunCommand struct{}

func (c *RunCommand) Execute(args []string) error {
	gofakeit.Seed(time.Now().UnixNano())
	go app.ProcessSignals()
	r := app.Router()
	if err := r.Run(":8080"); err != nil {
		return fmt.Errorf("cannot start the server on :8080 > %w", err)
	}
	return nil
}
