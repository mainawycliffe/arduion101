package main

import (
	"machine"
	"time"
)

func main() {

	led := machine.Pin(13)

	// same as var led = machine.LED
	led.Configure(machine.PinConfig{1})
	// 1 = output mode, 0 = input mode

	led_switch := false // led switch

	for {
		led.Set(led_switch)
		led_switch = !led_switch // invert switch
		if led_switch {
			delay(1000)
		} else {
			delay(500)
		}
	}
}

func delay(t int64) { // time delay function
	time.Sleep(time.Duration(1000000 * t))
}
