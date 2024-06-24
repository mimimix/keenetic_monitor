package userCommands

import (
	"go.uber.org/fx"
)

var NewFxUserCommands = fx.Module("bot.adminCommands",
	fx.Invoke(
		handlerGetID,
	),
)
