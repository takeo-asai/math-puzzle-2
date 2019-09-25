package q2

import (
	"math"
)

// Main :
func Main() float64 {
	s := 1
	t := 17
	N := 29

	n := (t - s)
	cnt := math.Pow(2, float64(n-1)) + math.Pow(2, float64(N-n-1)) - 1

	return cnt
}
