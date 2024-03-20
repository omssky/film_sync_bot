package bot

import (
	"film_sync_bot/internal/config"
	"film_sync_bot/internal/handlers"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"log"
	"time"
)

func StartBot() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  conf.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	bot.Use(middleware.Logger())

	bot.Handle("/start", handlers.StartHandler)

	log.Println("Bot " + bot.Me.Username + " started!")
	bot.Start()
}
