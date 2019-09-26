package q10

// Ceil : ceil(10, 3) == 4
func Ceil(x, y int) int {
	ceil := 0
	if x%y > 0 {
		ceil = 1
	}
	return x/y + ceil
}
