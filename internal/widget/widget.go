package widget

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
