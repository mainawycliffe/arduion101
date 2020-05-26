package main

import (
	"machine"
	"time"
)

func main() {

	// use pin 13, which is connected to the built-in LED
	var led machine.Pin = machine.Pin(13)

	// set the pin to output mode
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.High()                         // pull high (output power, light up led)
		time.Sleep(time.Millisecond * 500) // wait 500 milliseconds
		led.Low()                          // pull low (cut off power)
		time.Sleep(time.Millisecond * 500) // wait 500 milliseconds
	}
}
