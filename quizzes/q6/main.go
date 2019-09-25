package q6

// Main : Solve
func Main(n int) int {
	cnt := 0
	rect := Rectangle()
	for x := 1; x <= n; x++ {
		for y := 1; y <= x; y++ {
			if rect.CountSquares(x, y) == 20 {
				cnt++
			}
		}
	}
	return cnt
}
