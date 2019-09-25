package q5

// Main : Solve Problem
func Main(target int) int {
	// generate PascalTriangle recursively
	n := []int{1}
	tri := PascalTriangle()
	for i := 1; i <= target; i++ {
		n = tri.NextLine(n)
	}

	numberOfCash := 0
	for _, cash := range n {
		numberOfCash += CashCounter(cash).Count()
	}
	return numberOfCash
}
