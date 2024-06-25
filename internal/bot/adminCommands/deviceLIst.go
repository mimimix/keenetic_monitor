package adminCommands

import (
	"encoding/json"
	"fmt"
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
	getDevicesKB := &telebot.ReplyMarkup{}
	marshal, _ := json.Marshal([]any{"1"})
	getDevicesKB.Inline(
		getDevicesKB.Row(getDevicesKB.Data("Получить", "listDevices", string(marshal))),
	)
	fmt.Println(getDevicesKB)
	group.Handle("/btn", func(c telebot.Context) error {
		return c.Send("🌐 Список устройств", getDevicesKB)
	})

	sendDeviceListCommand := createDeviceListCommand(router)
	group.Handle("/devices", sendDeviceListCommand)
	//group.Handle(&telebot.Btn{Unique: "listDevices"}, sendDeviceListCommand)
	group.Handle(&telebot.Btn{Unique: "listDevices"}, func(c telebot.Context) error {
		fmt.Println(c.Data())
		return nil
	})
}
