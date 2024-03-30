package bot

import (
	"film_sync_bot/internal/config"
	"film_sync_bot/internal/handlers"
	"fmt"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"log"
	"time"
)

func StartBot() (*telebot.Bot, error) {
	conf, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  conf.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}

	bot.Use(middleware.Logger())

	bot.Handle("/start", handlers.StartHandler)

	log.Println("Bot " + bot.Me.Username + " started!")

	return bot, nil
}
