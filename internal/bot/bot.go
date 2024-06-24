package bot

import (
	"context"
	"go.uber.org/fx"
	"gopkg.in/telebot.v3"
	"keeneticmonitor/internal/bot/adminCommands"
	"keeneticmonitor/internal/bot/userCommands"
	"keeneticmonitor/pkg/config"
	"time"
)

func NewBot(config *config.AppConfig, lc fx.Lifecycle) *telebot.Bot {
	b, err := telebot.NewBot(telebot.Settings{
		Token:  config.TelegramToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		panic(err)
		return nil
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go b.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			b.Stop()
			return nil
		},
	})

	return b
}

var NewFxBot = fx.Module("bot",
	fx.Provide(
		NewBot,
	),
	adminCommands.NewFxAdminCommands,
	userCommands.NewFxUserCommands,
)
