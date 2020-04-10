package main

import (
	"log"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	adapter := raspi.NewAdaptor()
	led := gpio.NewLedDriver(adapter, "18")

	work := func() {
		gobot.Every(5*time.Second, func() {
			err := led.Toggle()
			if err != nil {
				log.Fatalf("An error occurred turning on the LED!")
			}
		})
	}

	robot := gobot.NewRobot("snapshot",
		[]gobot.Connection{adapter},
		[]gobot.Device{led},
		work,
	)

	err := robot.Start()
	if err != nil {
		log.Fatalf("Could not start the damn robot")
	}
}
