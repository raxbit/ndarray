package ndarray

import (
	"math"
	"math/rand"
)

func RandomInCircle(max float64) *Vec {
	t := 2 * math.Pi * rand.Float64()
	u := rand.Float64() + rand.Float64()
	r := u
	if u > 1 {
		r = 2 - u
	}
	r *= max

	return V(r*math.Cos(t), r*math.Sin(t), 0)
}
