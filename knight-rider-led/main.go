package main

import (
	"machine"
	"time"
)

const max_led_num = 5 // number of leds

func main() {

	var leds = [max_led_num]machine.Pin{
		machine.Pin(2),
		machine.Pin(3),
		machine.Pin(4),
		machine.Pin(5),
		machine.Pin(6),
	} // array of pins

	for i := 0; i < len(leds); i++ { // setup leds
		leds[i].Configure(machine.PinConfig{1})
	}

	for {

		// blink one by one forward
		for i := 0; i < len(leds); i++ {
			leds[i].High()
			delay(75)
			leds[i].Low()
		}

		// blink one by one backward
		for i := len(leds) - 1; i >= 0; i-- {
			leds[i].High()
			delay(75)
			leds[i].Low()
		}
	}
}

func delay(t int64) {
	time.Sleep(time.Duration(1000000 * t))
}
