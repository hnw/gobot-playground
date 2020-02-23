package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	board := raspi.NewAdaptor()
	bme280 := i2c.NewBME280Driver(board)
	tsl2561 := i2c.NewTSL2561Driver(board, i2c.WithTSL2561Gain16X, i2c.WithAddress(0x29))
	done := make(chan bool, 1)

	work := func() {

		t, _ := bme280.Temperature()
		fmt.Printf("Temperature: %v C\n", t)

		p, _ := bme280.Pressure()
		fmt.Printf("Pressure: %v hPa\n", p/100.0)

		h, _ := bme280.Humidity()
		fmt.Printf("Humidity: %v %%\n", h)

		broadband, ir, err := tsl2561.GetLuminocity()

		if err != nil {
			fmt.Println("Err:", err)
		} else {
			light := tsl2561.CalculateLux(broadband, ir)
			fmt.Printf("BB: %v, IR: %v, Lux: %v\n", broadband, ir, light)
		}
		done <- true
	}

	robot := gobot.NewRobot("bme280bot",
		[]gobot.Connection{board},
		[]gobot.Device{bme280, tsl2561},
		work,
	)

	err := robot.Start(false)
	if err != nil {
		return
	}
	<-done
	robot.Stop()
}
