package adminCommands

import "github.com/mimimix/go-keenetic-api"

func createDeviceName(device keenetic.Device) string {
	prefix := "⚾️"
	if device.Active {
		prefix = "🥎"
	}
	name := device.Mac
	if device.Name != "" {
		name = device.Name
	}
	return prefix + " " + name
}
