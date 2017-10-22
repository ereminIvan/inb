package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/iveronanomi/inbogo/app"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	signals := make(chan os.Signal, 1)
	shutdown := make(chan struct{})

	defer close(signals)
	defer close(shutdown)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	a := app.New(shutdown)

	go func(){
		<- signals
		shutdown <- struct{}{}
	}()

	a.Run()
}
