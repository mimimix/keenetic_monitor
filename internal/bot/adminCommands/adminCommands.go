package adminCommands

import (
	"go.uber.org/fx"
	"gopkg.in/telebot.v3"
)

func newAdminGroup(bot *telebot.Bot) *telebot.Group {
	adminGroup := bot.Group()
	return adminGroup
	//adminBot.Use(middleware.Whitelist(a.config.MainAdminId...))
}

var NewFxAdminCommands = fx.Module("bot.adminCommands",
	fx.Provide(
		fx.Private,
		newAdminGroup,
	),
	fx.Invoke(
		handlerDeviceList,
		handlerPoller,
	),
)
