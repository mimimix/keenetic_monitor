package adminCommands

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/mimimix/go-keenetic-api"
	"gopkg.in/telebot.v3"
	"sort"
	"strings"
	"time"
)

func createDeviceListCommand(router *keenetic.Keenetic, poller *keenetic.Poller) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		devices, err := router.DeviceList()
		if err != nil {
			return c.Send(err.Error())
		}
		sort.Slice(*devices, func(i, j int) bool {
			return (*devices)[j].Active != (*devices)[i].Active
		})
		var textDevices []string
		for _, device := range *devices {
			var eventTime *time.Time
			if !device.Active {
				if eventTime == nil {
					eventTime = nil
				} else {
					lastTime, isExists := poller.GetLastOnline(device.Mac)
					if isExists {
						eventTime = &lastTime
					} else {
						eventTime = nil
					}
				}
			} else {
				eventSince := time.Now().Add(-time.Duration(device.Uptime) * time.Second)
				eventTime = &eventSince
			}
			uptimeStr := ""
			if eventTime != nil {
				uptimeStr = " - " + humanize.Time(*eventTime)
			}
			textDevices = append(textDevices, createDeviceName(device)+uptimeStr)
		}
		return c.Send("🌐 Список устройств:\n\n" + strings.Join(textDevices, "\n"))
	}
}

func handlerDeviceList(group *telebot.Group, bot *telebot.Bot, router *keenetic.Keenetic, poller *keenetic.Poller) {
	getDevicesKB := &telebot.ReplyMarkup{}
	getDevicesKB.Inline(
		getDevicesKB.Row(getDevicesKB.Data("Получить", "listDevices")),
	)
	fmt.Println(getDevicesKB)
	group.Handle("/btn", func(c telebot.Context) error {
		return c.Send("🌐 Список устройств", getDevicesKB)
	})

	sendDeviceListCommand := createDeviceListCommand(router, poller)
	group.Handle("/devices", sendDeviceListCommand)
	group.Handle(&telebot.Btn{Unique: "listDevices"}, sendDeviceListCommand)
}
