package main

import (
	"go.uber.org/fx"
	"keeneticmonitor/internal/app"
)

func main() {
	fx.New(app.App).Run()
}
