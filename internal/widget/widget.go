package widget

import (
	"encoding/json"
	"fmt"

	"github.com/daveszczesny/project-conos/internal/constants"
)

type WidgetState struct {
	State string
}

type WidgetObject struct {
	Widget Widget `json: "widget"`
}

type Widget struct {
	Type string `json: "type"`
	Id   uint8  `json: "id"`
	Icon struct {
		Src  string `json: "src"`
		Type string `json: "type"`
	} `json: "icon"`

	Action []struct {
		Function  string `json: "function"`
		Frequency int    `json: "frequency"`
	} `json: "action"`

	RefreshRate int `json: "refreshRate"`
	ExpiryTime  int `json: "expiryTime"`
	Position    struct {
		X int `json: "x"`
		Y int `json: "y"`
	} `json: "position"`
	Size struct {
		Width  int `json: "w"`
		Height int `json: "h"`
	} `json: "size"`
	State constants.WidgetState
}

func CreateWidget(jsonData string) *Widget {
	var myWidget Widget
	err := json.Unmarshal([]byte(jsonData), &myWidget)
	if err != nil {
		fmt.Println("Error. Could not parse json file")
		return nil
	}
	if myWidget.Type != "Widget" {
		fmt.Println("Error. Type mismatch. Expected Widget type of Widget, but got ", myWidget.Type)
		return nil
	}
	myWidget.State = constants.Inactive

	return &myWidget
}

/*
* Function to create a Widget Folder
* A widget folder contains a list of widgets
* Each element of the list contains widget data (position and size) and the widget
 */
func CreateWidgetsFromJson(jsonData string) ([]WidgetObject, error) {
	var data struct {
		Widgets []WidgetObject `json "Widgets"`
	}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return nil, fmt.Errorf("Error parsing json data into WidgetObjects: %v", err)
	}

	for i := range data.Widgets {
		data.Widgets[i].Widget.State = constants.Inactive
	}

	return data.Widgets, nil
}
