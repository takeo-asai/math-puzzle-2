package q12_test

import (
	"fmt"
	"testing"

	"puzzle/quizzes/q12"
)

func Test_Solve(t *testing.T) {
	fmt.Printf("n = 1, %v\n", q12.Solve(1))
	fmt.Printf("n = 2, %v\n", q12.Solve(2))
	fmt.Printf("n = 11, %v\n", q12.Solve(11))
}
