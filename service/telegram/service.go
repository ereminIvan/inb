package telegram

import (
	"log"
	"time"

	api "github.com/iveronanomi/telegram-bot-api"

	"github.com/iveronanomi/inbogo/model"
)

type service struct {
	client *api.BotAPI
	config model.TelegramConfig
	updateConfig api.UpdateConfig
}

func New(config model.TelegramConfig) *service {
	bot, err := api.NewBotAPI(config.Token)
	if err != nil {
		panic(err)
	}
	bot.Debug = config.DebugEnabled
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &service{
		client: bot,
		config: config,
		updateConfig: api.UpdateConfig{
			Offset: config.Offset,
			Limit: config.Limit,
			Timeout: config.Timeout,
		},
	}
}

func (s *service) Run(interrupt chan struct{}) {

	log.Print("Run")
	updatesChan := make(chan api.Update, 100)
	s.updateConfig.Timeout = 60

	go func() {
		for {
			log.Print("rutine: get updates")
			select {
				case <-interrupt:
					log.Print("Interrupt recived")
					return
				default:
			}
			updates, err := s.client.GetUpdates(s.updateConfig)
			if err != nil {
				log.Printf("Failed to get updates, retrying in 3 seconds... %v", err)
				time.Sleep(time.Second * 3)
				continue
			}
			//update config for update
			for _, update := range updates {
				if update.UpdateID >= s.updateConfig.Offset {
					s.updateConfig.Offset = update.UpdateID + 1
					updatesChan <- update
				}
			}
		}
	}()

	go func() {
		log.Print("rutine: read updates")
		select {
			case <-interrupt:
				log.Print("Interrupt recived")
				return
			default:
		}
		select {
			case update := <-updatesChan: {
				if update.Message != nil {
					log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

					msg := api.NewMessage(update.Message.Chat.ID, update.Message.Text)
					msg.ReplyToMessageID = update.Message.MessageID
				}
			}
		}
	}()
	<-interrupt
	log.Print("TG interrupt recived")
}
