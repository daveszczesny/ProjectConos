package main

import (
	"fmt"

	"github.com/daveszczesny/project-conos/internal/widget"
)

func main() {

	jsonData := `
	{
		"Widgets": [
			{
				"widget": {
					"type": "Widget",
					"id": 123,
					"icon": {
						"src": "image",
						"type": "PNG"
					},
					"action": [
						{
							"function": "getDateAndTime",
							"frequency": 100
						}
					],
					"refreshRate": 100,
					"expiryTime": 10000,
					"position": {
						"x": 0,
						"y": 0
					},
					"size": {
						"w": 500,
						"h": 500
					}
				}
			},
			{
				"widget": {
					"type": "Widget",
					"id": 222,
					"icon": {
						"src": "different",
						"type": "PNG"
					},
					"action": [
						{
							"function": "getSpotifySong",
							"frequency": 200
						}
					],
					"refreshRate": 100,
					"expiryTime": 10000,
					"position": {
						"x": 100,
						"y": 100
					},
					"size": {
						"w": 500,
						"h": 500
					}
				}
			}
		]
	}
	`

	widgetObjects, err := widget.CreateWidgetsFromJson(jsonData)
	if err != nil {
		fmt.Printf("Error creating WidgetObjects: %v\n", err)
		return
	}

	for _, obj := range widgetObjects {
		fmt.Println(obj.Widget.State)
	}
}
