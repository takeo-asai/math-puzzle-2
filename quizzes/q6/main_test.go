package q6_test

import (
	"testing"

	"puzzle/quizzes/q6"
)

func Test_Solve(t *testing.T) {
	n := q6.Main(1000)
	t.Logf("the result is %v", n)
}
