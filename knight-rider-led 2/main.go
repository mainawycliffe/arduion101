package main

import (
	"machine"
	"time"
)

const (
	max_led_num   = 5
	start_led_pin = 2
)

func main() {

	delay := func(t int64) { // implicit time delay function
		time.Sleep(time.Duration(1000000 * t))
	}

	var leds [max_led_num]machine.Pin // array of leds
	index, delta := 0, 1              // blink index and direction

	for i, _ := range leds { // setup leds
		leds[i] = machine.Pin(i + start_led_pin)
		leds[i].Configure(machine.PinConfig{1})
	}

	for {

		// turn one the led on index and turn off the rest
		for i, led := range leds {
			led.Set(index == i)
		}

		index += delta // move index

		// invert index's moving direction at each end
		if index == 0 || index == len(leds)-1 {
			delta *= -1
		}

		delay(100)
	}
}
