package q3_test

import (
	"testing"

	"puzzle/quizzes/q3"
)

func Test_Solve(t *testing.T) {
	cnt := q3.Main()
	t.Logf("The result is %v", cnt)
	if cnt != 93 {
		t.Errorf("The result must be 93, but %v", cnt)
	}
}
