package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	adapter := raspi.NewAdaptor()
	led := gpio.NewLedDriver(adapter, "18")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("snapshot",
		[]gobot.Connection{adapter},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
