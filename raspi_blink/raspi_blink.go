package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	greenLed := gpio.NewLedDriver(r, "11")  // 11: GPIO 17
	yellowLed := gpio.NewLedDriver(r, "12") // 12: GPIO 18
	blueLed := gpio.NewLedDriver(r, "15")   // 15: GPIO 22
	whiteLed := gpio.NewLedDriver(r, "13")  // 13: GPIO 27

	work := func() {
		gobot.Every(1*time.Second, func() {
			greenLed.Toggle()
		})
		gobot.Every(2*time.Second, func() {
			yellowLed.Toggle()
		})
		gobot.Every(4*time.Second, func() {
			blueLed.Toggle()
		})
		gobot.Every(8*time.Second, func() {
			whiteLed.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{greenLed, yellowLed, blueLed, whiteLed},
		work,
	)

	robot.Start()
}
