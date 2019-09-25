package q6

// IRectangle :
type IRectangle interface{ CountSquares(x, y int) int }

// Rectangle : constructor for IRectangle
func Rectangle() IRectangle {
	return &rect{map[int]int{}}
}

type rect struct {
	cache map[int]int
}

func (r *rect) CountSquares(x, y int) int {
	if x > 1000 {
		panic("'size > 1000' is not supported")
	}

	if y > x {
		return r.CountSquares(y, x)
	}
	// If there is a cache, use it
	v, ok := r.cache[x*1000+y]
	if ok {
		return v
	}
	if x == 1 || x == y {
		return 1
	}
	n := 1 + r.CountSquares(x-y, y)
	r.cache[x*1000+y] = n
	return n
}
