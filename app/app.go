package app

import (
	"log"
	"sync"

	"github.com/iveronanomi/inbogo/model"
	"github.com/iveronanomi/inbogo/service/telegram"
)

type ITGService interface {
	Run(interrupt chan struct{})
}

type application struct {
	tgService ITGService
	interrupt chan struct{}
	sync.Once
}

func New(interrupt chan struct{}) *application {
	return &application{
		tgService: telegram.New(model.TelegramConfig{
			Token: "461634935:AAFgb3fJR-MenuHSiW9gunndkZ_oOEPhebY",
		}),
		interrupt: interrupt,
	}
}

func (a *application) Run() {
	log.Print("Run")
	go func(){
		a.tgService.Run(a.interrupt)
	}()
	<-a.interrupt
	a.stop()
}

func (a *application) stop() {
	log.Print("Stop")
}