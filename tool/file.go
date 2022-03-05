package tool

import "math"

func ByteToM(size int64) float64 {
	pow := math.Pow(1024, 2)
	result := float64(size)/ pow
	return math.Round((result*100)/100)
}
