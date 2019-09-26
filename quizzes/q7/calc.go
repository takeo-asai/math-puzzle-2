package q7

func nPr(n, r int) int64 {
	var result int64 = 1
	for i := 0; i < r; i++ {
		result *= int64(n - i)
	}
	return result
}

func F(n, k int) int64 {
	return int64(n-k) * nPr(n, k-1)
}
