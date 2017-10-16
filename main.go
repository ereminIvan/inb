package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ereminIvan/inbogo/app"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	shutdown := make(chan struct{})

	a := app.New()

	go func(){
		<- signals
		a.Stop()
		shutdown <- struct{}{}
	}()

	a.Run(shutdown)

	close(signals)
	close(shutdown)
}
