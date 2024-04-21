package main

import (
	"machine"

	"github.com/daveszczesny/project-cronos/internal/json"
)

var json_data = `{
	"icon": "img",
	"action": "func",
	"refreshRate": 100,
	"expiryTime": 50,
	"position": [10; 40],
	"size": [10; 10]
}`

var led = machine.Pin(machine.D6)
var btn = machine.Pin(machine.D8)
var widget, _ = json.JSONDecoder(json_data)

func configurePins() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	btn.Configure(machine.PinConfig{Mode: machine.PinInput})
}

func main() {
	configurePins()

	for {
		if btn.Get() == true {
			led.Set(true)
		} else {
			led.Set(false)
		}
	}
}
