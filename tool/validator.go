package tool

var VideoFormats = map[string]int{
	"mp4":0,
	"MOV":1,
	"WMV":2,
	"AVI":3,
	"MKV":4,
	"MPEG-2":5,
}

func IsVideo(fileFormat string) bool  {
	_, exists  := VideoFormats[fileFormat]
	return exists
}
