package q11_test

import (
	"testing"

	"puzzle/quizzes/q11"
)

func Test_Solve(t *testing.T) {
	t.Logf("n=%v, result=%v", 3, q11.Vote(3))
	t.Logf("n=%v, result=%v", 7, q11.Vote(7))
}
