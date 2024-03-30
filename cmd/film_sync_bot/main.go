package main

import (
	"film_sync_bot/internal/bot"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	telebot, err := bot.StartBot()
	if err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		telebot.Start()
	}()

	<-quit

	telebot.Stop()
}
