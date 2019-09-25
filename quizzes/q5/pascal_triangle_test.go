package q5_test

import (
	"testing"

	"puzzle/quizzes/q5"
)

func Test_CalculateNextTriangle(t *testing.T) {
	tri := q5.PascalTriangle()
	n1 := []int{1, 1}
	n2 := tri.NextLine(n1)
	n3 := tri.NextLine(n2)
	n4 := tri.NextLine(n3)
	n5 := tri.NextLine(n4)
	t.Logf("value: %v", n2)
	t.Logf("value: %v", n3)
	t.Logf("value: %v", n4)
	t.Logf("value: %v", n5)
}
