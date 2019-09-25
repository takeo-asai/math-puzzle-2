package q4_test

import (
	"testing"

	"puzzle/quizzes/q4"
)

func Test_Solve(t *testing.T) {
	cnt := q4.Main()
	t.Logf("the result is %v", cnt)
}
