package keenetic

import (
	"github.com/mimimix/go-keenetic-api"
	"keeneticmonitor/pkg/config"
)

func NewKeeneticClient(config *config.AppConfig) *keenetic.Keenetic {
	return keenetic.NewKeenetic(config.KeeneticUsername, config.KeeneticPassword, config.KeeneticBaseUrl)
}
