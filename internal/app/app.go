package app

import (
	"go.uber.org/fx"
	"keeneticmonitor/internal/bot"
	"keeneticmonitor/pkg/config"
	"keeneticmonitor/pkg/keenetic"
)

var App = fx.Options(
	fx.Provide(
		config.NewConfig,
		keenetic.NewKeeneticClient,
	),
	fx.Invoke(),
	bot.NewFxBot,
)
