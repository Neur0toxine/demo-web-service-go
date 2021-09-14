package internal

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// ProcessSignals will react to app termination. It doesn't do much for now, but it can be useful later.
// For example, we can close active DB or WebSocket connections here, or perform any other actions
// related to the graceful shutdown.
func ProcessSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for {
		<-c
		log.Println("Quitting...")
		os.Exit(0)
	}
}
