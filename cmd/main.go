package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/arthurhzna/go-clean-architecture/internal/bootstrap"
)

func main() {
	app := bootstrap.NewApplication()

	go app.HttpServer.Start()

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-quit

	app.HttpServer.Shutdown()
}
