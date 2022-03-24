package tool

import (
	"math"
	"strings"
)

func ByteToM(size int64) float64 {
	pow := math.Pow(1024, 2)
	result := float64(size)/ pow
	return math.Round((result*100)/100)
}

func ParseFileName(fileName string) (string, string) {
	split := strings.Split(fileName, ".")
	name := split[0]
	format := split[1]
	return name, format
}
