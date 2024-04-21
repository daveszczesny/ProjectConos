package widget

/*
* Widget: 	Denotes a widget that will be displayed
*			on screen, contains
			- icon: image in PNG format
			- action: function it calls
			- refresh rate: time between widget refresh
			- expiry time: time left before widget dies
			- position: position on screen
			- size: size on screen
*/
type Widget struct {
	Icon        string
	Action      string
	RefreshRate uint8
	ExpiryTime  uint16
	Position    struct {
		X uint8
		Y uint8
	}
	Size struct {
		W uint16
		H uint16
	}
}
