package json

import (
	"strconv"
	"strings"

	"github.com/daveszczesny/project-cronos/internal/widget"
)

func JSONDecoder(jsonData string) (*widget.Widget, string) {
	var widget widget.Widget

	jsonData = strings.ReplaceAll(jsonData, "\n", "")
	jsonData = strings.TrimSpace(jsonData)

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
			return nil, "error message, start error"
		}

		start += len(key)
		end := strings.Index(jsonData[start:], ",")
		if end == -1 {
			end = strings.Index(jsonData[start:], "}")
		}

		if end == -1 {
			return nil, "error message, end error"
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
				return nil, "error message, refresh rate error"
			}
			widget.RefreshRate = uint8(refreshRate)
		case "ExpiryTime":
			expiryTime, err := strconv.ParseUint(valueStr, 10, 16)
			if err != nil {
				return nil, "error message, expiry time error"
			}
			widget.ExpiryTime = uint16(expiryTime)
		case "Position":
			xy := strings.Split(valueStr, ";")
			if len(xy) != 2 {
				return nil, "Error message, two parts not found in position"
			}
			x, err := strconv.ParseUint(strings.TrimSpace(strings.ReplaceAll(xy[0], "[", "")), 10, 8)
			if err != nil {
				return nil, "Error message, x is incorrect"
			}
			y, err := strconv.ParseUint(strings.TrimSpace(strings.ReplaceAll(xy[1], "]", "")), 10, 8)
			if err != nil {
				return nil, "Error message, y is incorrect"
			}
			widget.Position.X = uint8(x)
			widget.Position.Y = uint8(y)
		case "Size":
			wh := strings.Split(valueStr, ";")
			if len(wh) != 2 {
				return nil, "Error message, two parts not found in size"
			}
			w, err := strconv.ParseUint(strings.TrimSpace(strings.ReplaceAll(wh[0], "[", "")), 10, 16)
			if err != nil {
				return nil, "Error message, width is incorrect"
			}
			h, err := strconv.ParseUint(strings.TrimSpace(strings.ReplaceAll(wh[1], "]", "")), 10, 16)
			if err != nil {
				return nil, "Error message, height is incorrect"
			}
			widget.Size.W = uint16(w)
			widget.Size.H = uint16(h)

		}

	}

	return &widget, ""
}
