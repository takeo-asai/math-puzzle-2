package q12

import (
	"math"
)

// Solve :
func Solve(n int64) int64 {
	pi := Pi(n)
	var i int64
	for i = 1; true; i++ {
		a := int(float64(pi.Value()*i) / math.Pow(10, float64(pi.N())))
		b := float64((pi.Value()+1)*i) / math.Pow(10, float64(pi.N()))
		if a != int(b) && float64(int(b)) != b { // a <= n < b: bが整数だと比較に困る
			return i
		}
	}
	return -1
}
