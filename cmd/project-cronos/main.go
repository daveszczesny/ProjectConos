package main

import (
	"machine"

	"tinygo.org/x/drivers/gc9a01"

	"github.com/daveszczesny/project-cronos/internal/constants"
	"github.com/daveszczesny/project-cronos/internal/json"
)

// static variables
var led1 = machine.Pin(constants.LED_PIN_1)
var led2 = machine.Pin(constants.LED_PIN_2)
var btn = machine.Pin(constants.BTN_PIN)
var widget, _ = json.JSONDecoder(constants.JSON_DATA)

// Display variables
var spi = machine.SPI0
var display gc9a01.Device = gc9a01.New(spi,
	constants.RST_PIN,
	constants.DC_PIN,
	constants.CS_PIN,
	constants.BL_PIN)

/*
*	Configures the pin config for all pins
 */
func configurePins() {
	led1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	btn.Configure(machine.PinConfig{Mode: machine.PinInput})
	// display.Configure(gc9a01.Config{...})
}

/*
*	Entry point of program
 */
func main() {
	configurePins()

	for {

		// Button actions to trigger LED
		if btn.Get() == true {
			led1.Set(true)
			led2.Set(true)
		} else {
			led1.Set(false)
			led2.Set(false)
		}
	}
}
