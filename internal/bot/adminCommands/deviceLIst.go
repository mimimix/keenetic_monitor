package adminCommands

import (
	"github.com/mimimix/go-keenetic-api"
	"gopkg.in/telebot.v3"
	"sort"
	"strings"
)

func createDeviceListCommand(router *keenetic.Keenetic) func(c telebot.Context) error {
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
			textDevices = append(textDevices, createDeviceName(device))
		}
		return c.Send("🌐 Список устройств:\n\n" + strings.Join(textDevices, "\n"))
	}
}

func handlerDeviceList(group *telebot.Group, router *keenetic.Keenetic) {
	println("moe")
	getDevicesKB := &telebot.ReplyMarkup{}
	getDevicesBtn := getDevicesKB.Data("Получить", "listDevices")
	getDevicesKB.Inline(
		getDevicesKB.Row(getDevicesBtn),
	)
	group.Handle("/btn", func(c telebot.Context) error {
		return c.Send("🌐 Список устройств", getDevicesKB)
	})

	sendDeviceListCommand := createDeviceListCommand(router)
	group.Handle("/devices", sendDeviceListCommand)
	group.Handle(&getDevicesBtn, sendDeviceListCommand)
}
