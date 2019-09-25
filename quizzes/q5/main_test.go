package q5_test

import (
	"testing"

	"puzzle/quizzes/q5"
)

func Test_Solve(t *testing.T) {
	numOfCash := q5.Main(45)
	t.Logf("the result is %v", numOfCash)
	if numOfCash != 3518437540 {
		t.Errorf("n=45, the result must be 3518437540, but %v", numOfCash)
	}
}
