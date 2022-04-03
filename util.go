package main

import "math/rand"

func Random(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}
