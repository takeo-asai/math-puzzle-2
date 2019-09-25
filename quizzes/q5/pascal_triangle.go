package q5

// IPascalTriangle :
type IPascalTriangle interface{ NextLine(line []int) []int }

// PascalTriangle : constructor for IPascalTriangle
func PascalTriangle() IPascalTriangle {
	return &pascal{}
}

type pascal struct{}

func (p *pascal) NextLine(line []int) []int {
	var nextLine = []int{1}
	for i := 0; i < len(line)-1; i++ {
		nextLine = append(nextLine, line[i]+line[i+1])
	}
	return append(nextLine, 1)
}
