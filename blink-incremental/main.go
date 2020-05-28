package main

import (
	"machine"
	"time"
)

const (
	max_led_num   = 6
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

		for x := 0; x <= index; x++ {
			leds[x].Set(true)
			delay(150)
		}
		delay(500)
		for x := index; x >= 0; x-- {
			leds[x].Set(false)
			delay(150)
		}
		// // turn one the led on index and turn off the rest
		// for i, led := range leds {
		// 	led.Set(i <= index)
		// }

		index += delta

		if index == 0 || index == len(leds)-1 {
			delta *= -1
		}

		delay(200)
	}
}
