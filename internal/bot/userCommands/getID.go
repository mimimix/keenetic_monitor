package userCommands

import (
	"fmt"
	"gopkg.in/telebot.v3"
)

func handlerGetID(bot *telebot.Bot) {
	bot.Handle("/id", func(c telebot.Context) error {
		return c.Send(fmt.Sprintf("💬 ChatID: %d\n👤 UserID: %d", c.Chat().ID, c.Sender().ID))
	})
}
