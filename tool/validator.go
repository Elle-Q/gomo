package tool

import "strings"

var VideoFormats = map[string]int{
	"video/mp4": 0,
	"mp4":       0,
	"MOV":       1,
	"WMV":       2,
	"AVI":       3,
	"MKV":       4,
	"MPEG-2":    5,
}

func IsVideo(fileFormat string) bool {
	fileFormat = strings.ToLower(fileFormat)
	_, exists := VideoFormats[fileFormat]
	return exists
}
