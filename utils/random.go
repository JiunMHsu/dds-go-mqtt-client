package utils

import (
	"math"
	"math/rand/v2"
)

func RandomFloat(min, max float64) float64 {
	return math.Round((rand.Float64()*(max-min)+min)*100) / 100
}
