package handlers

import (
	"gopkg.in/telebot.v3"
)

func StartHandler(ctx telebot.Context) error {
	return ctx.Reply("Hello telegram!")
}
