package constants

type WidgetState string

const (
	Inactive     WidgetState = "Inactive"
	Active       WidgetState = "Active"
	Initializing WidgetState = "Initializing"
	Ready        WidgetState = "Ready"
)

type Position struct {
	X int
	Y int
}

type Size struct {
	WIDTH  int
	HEIGHT int
}
