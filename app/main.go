package main

import (
	"log"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/rpi"
)

func main() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	t := time.NewTicker(1 * time.Second)
	for l := gpio.Low; ; l = !l {
		if err := rpi.P1_18.Out(l); err != nil {
			log.Fatal(err)
		}

		if err := rpi.P1_40.Out(l); err != nil {
			log.Fatal(err)
		}
		<-t.C
	}

}
