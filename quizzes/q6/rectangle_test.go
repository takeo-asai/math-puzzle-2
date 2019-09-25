package q6_test

import (
	"testing"

	"puzzle/quizzes/q6"
)

func Test_CalcRectangles(t *testing.T) {
	rect := q6.Rectangle()
	n := rect.CountSquares(10, 2)
	t.Logf("count %v", n)
}
