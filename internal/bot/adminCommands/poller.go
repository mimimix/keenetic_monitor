package adminCommands

import (
	"fmt"
	"github.com/mimimix/go-keenetic-api"
	"github.com/zmwangx/debounce"
	"gopkg.in/telebot.v3"
	"keeneticmonitor/pkg/config"
	"strings"
	"time"
)

func handlerPoller(bot *telebot.Bot, poller *keenetic.Poller, config *config.AppConfig) {
	if config.PollingIsEnabled && config.PollingChatId != 1 {
		var sendQueue []string

		debounceSend, _ := debounce.Debounce(func() {
			_, _ = bot.Send(telebot.ChatID(config.PollingChatId), strings.Join(sendQueue, "\n"))
			sendQueue = []string{}
		}, 200*time.Millisecond, debounce.WithMaxWait(time.Second))

		go func() {
			for {
				event := <-poller.Channel
				fmt.Println(event)
				sendQueue = append(sendQueue, createDeviceName(*event.Device))
				debounceSend()
			}
			//_, _ = Bot.Send(telebot.ChatID(config.PollingChatId), CreateDeviceName(*event.Device))
		}()
		poller.Start()
	}
}
