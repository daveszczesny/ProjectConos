package main

import (
	"machine"

	"github.com/daveszczesny/project-cronos/internal/constants"
	"github.com/daveszczesny/project-cronos/internal/json"
)

// static variables
var led = machine.Pin(machine.D6)
var btn = machine.Pin(machine.D8)
var widget, _ = json.JSONDecoder(constants.JSON_DATA)

/*
*	Configures the pin config for all pins
 */
func configurePins() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	btn.Configure(machine.PinConfig{Mode: machine.PinInput})
}

/*
*	Entry point of program
* 	Program for loop
 */
func main() {
	configurePins()

	for {

		// Button actions to trigger LED
		if btn.Get() == true {
			led.Set(true)
		} else {
			led.Set(false)
		}
	}
}
