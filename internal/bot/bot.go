package bot

import (
	"context"
	"fmt"
	"github.com/mimimix/go-keenetic-api"
	"go.uber.org/fx"
	"gopkg.in/telebot.v3"
	"keeneticmonitor/internal/bot/adminCommands"
	"keeneticmonitor/internal/bot/userCommands"
	"keeneticmonitor/pkg/config"
	"time"
)

func NewBot(config *config.AppConfig, lc fx.Lifecycle) *telebot.Bot {
	if config.TelegramToken == "" {
		panic("TELEGRAM TOKEN IS EMPTY")
	}
	b, err := telebot.NewBot(telebot.Settings{
		Token:  config.TelegramToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		OnError: func(err error, c telebot.Context) {
			fmt.Println("TELEGRAM ERROR: ", err)
		},
		//ParseMode: telebot.ModeMarkdownV2,
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

func NewPoller(router *keenetic.Keenetic, config *config.AppConfig) *keenetic.Poller {
	if config.PollingInterval < 1 {
		config.PollingInterval = 1
	}
	return keenetic.NewPoller(router, config.PollingInterval)
}

var NewFxBot = fx.Module("bot",
	fx.Provide(
		NewBot,
	),
	fx.Provide(
		fx.Private,
		NewPoller,
	),
	adminCommands.NewFxAdminCommands,
	userCommands.NewFxUserCommands,
)
