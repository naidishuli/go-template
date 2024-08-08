package utils

import "math"

func RoundTo3DP(num float64) float64 {
	return math.Round(num*float64(1000)) / float64(1000)
}
