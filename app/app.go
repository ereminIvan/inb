package app

import (
	"log"

	"github.com/ereminIvan/inbogo/service"
)

type application struct {
	tgService interface{}
}

func New() *application {
	return &application{
		tgService: service.NewTG("61634935:AAFgb3fJR-MenuHSiW9gunndkZ_oOEPhebY"),
	}
}

func (a *application) Run(shutdown chan struct{}) {
	log.Print("Run")
	for {
		select {
		case <-shutdown:
			return
		default:
		}
	}
}

func (a *application) Stop() {
	log.Print("Stop")
}