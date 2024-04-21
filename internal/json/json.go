package json

import (
	"strconv"
	"strings"

	"github.com/daveszczesny/project-cronos/internal/widget"
)

/*
  - JSON decoder for Widget Object
    :return Widget Pointer, Error String
    :param Json String
*/
func JSONDecoder(jsonData string) (*widget.Widget, string) {
	var widget widget.Widget

	jsonData = strings.ReplaceAll(jsonData, "\n", "")
	jsonData = strings.TrimSpace(jsonData)

	// Format of json file, must contain each of these attributes
	fields := map[string]string{
		"\"icon\":":        "Icon",
		"\"action\":":      "Action",
		"\"refreshRate\":": "RefreshRate",
		"\"expiryTime\":":  "ExpiryTime",
		"\"position\":":    "Position",
		"\"size\":":        "Size",
	}

	for key, fieldName := range fields {
		start := strings.Index(jsonData, key)
		if start == -1 {
			return nil, "Error in JSON data!"
		}

		start += len(key)
		end := strings.Index(jsonData[start:], ",")
		if end == -1 {
			end = strings.Index(jsonData[start:], "}")
		}

		if end == -1 {
			return nil, "Error parsing JSON data!"
		}

		valueStr := strings.TrimSpace(jsonData[start : start+end])
		valueStr = strings.Trim(valueStr, ":")

		switch fieldName {
		case "Icon":
			widget.Icon = strings.Trim(valueStr, "\"")
		case "Action":
			widget.Action = strings.Trim(valueStr, "\"")
		case "RefreshRate":
			refreshRate, err := strconv.ParseUint(valueStr, 10, 8)
			if err != nil {
				return nil, "Error processing: " + string(valueStr) + " in Refresh Rate"
			}
			widget.RefreshRate = uint8(refreshRate)
		case "ExpiryTime":
			expiryTime, err := strconv.ParseUint(valueStr, 10, 16)
			if err != nil {
				return nil, "Error processing: " + string(valueStr) + " in Expiry Time"
			}
			widget.ExpiryTime = uint16(expiryTime)

		// Both Position and Size are arrays
		// For simplicity, within arrays semi-colons are used instead of commas
		case "Position":
			coords := strings.Split(valueStr, ";")
			if len(coords) != 2 {
				return nil, "Error processing Position. Coordinates length not equal to 2"
			}
			x, err := strconv.ParseUint(strings.TrimSpace(strings.ReplaceAll(coords[0], "[", "")), 10, 8)
			if err != nil {
				return nil, "Error processing x component of coordinates"
			}
			y, err := strconv.ParseUint(strings.TrimSpace(strings.ReplaceAll(coords[1], "]", "")), 10, 8)
			if err != nil {
				return nil, "Error processing y component of coordinates"
			}
			widget.Position.X = uint8(x)
			widget.Position.Y = uint8(y)
		case "Size":
			res := strings.Split(valueStr, ";")
			if len(res) != 2 {
				return nil, "Error message, two parts not found in size"
			}
			width, err := strconv.ParseUint(strings.TrimSpace(strings.ReplaceAll(res[0], "[", "")), 10, 16)
			if err != nil {
				return nil, "Error message, width is incorrect"
			}
			height, err := strconv.ParseUint(strings.TrimSpace(strings.ReplaceAll(res[1], "]", "")), 10, 16)
			if err != nil {
				return nil, "Error message, height is incorrect"
			}
			widget.Size.W = uint16(width)
			widget.Size.H = uint16(height)

		}
	}

	return &widget, ""
}
