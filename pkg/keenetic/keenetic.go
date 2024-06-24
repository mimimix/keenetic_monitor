package keenetic

import (
	"fmt"
	"github.com/mimimix/go-keenetic-api"
	"keeneticmonitor/pkg/config"
)

func NewKeeneticClient(config *config.AppConfig) *keenetic.Keenetic {
	fmt.Println(config.KeeneticUsername, config.KeeneticPassword, config.KeeneticBaseUrl)
	return keenetic.NewKeenetic(config.KeeneticUsername, config.KeeneticPassword, config.KeeneticBaseUrl)
}
