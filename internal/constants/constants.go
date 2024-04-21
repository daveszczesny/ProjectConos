package constants

import "machine"

var JSON_DATA string = `{
	"icon": "img",
	"action": "func",
	"refreshRate": 100,
	"expiryTime": 50,
	"position": [10; 40],
	"size": [10; 10]
}`

const (
	// PIN CONSTANTS
	LED_PIN_1 machine.Pin = machine.D6
	LED_PIN_2 machine.Pin = machine.D7
	BTN_PIN   machine.Pin = machine.D8
	// GC9A01 PIN CONSTANTS
	DC_PIN  machine.Pin = machine.D1
	RST_PIN machine.Pin = machine.D2
	CS_PIN  machine.Pin = machine.D3
	BL_PIN  machine.Pin = machine.D4
	// DISPLAY RESOLUTION
	DISPLAY_WIDTH  int16 = 240
	DISPLAY_HEIGHT int16 = 230
)
