package q5_test

import (
	"testing"

	"puzzle/quizzes/q5"
)

func Test_Check_CachCounter(t *testing.T) {
	cnt := 0
	n9 := []int{1, 9, 36, 84, 126, 126, 84, 36, 9, 1}
	for _, c := range n9 {
		cnt += q5.CashCounter(c).Count()
	}
	if cnt != 48 {
		t.Errorf("n=9, the result is 48, but %v", cnt)
	}
}
