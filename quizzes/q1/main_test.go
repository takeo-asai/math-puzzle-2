package q1_test

import (
	"testing"

	"puzzle/quizzes/q1"
)

func Test_q1(t *testing.T) {
	cnt := q1.Main(4)
	t.Logf("total count is %v", cnt)
	if cnt != 5100 {
		t.Errorf("q1 result is 5100, but %v", cnt)
	}
}
