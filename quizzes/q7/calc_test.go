package q7_test

import (
	"testing"

	"puzzle/quizzes/q7"
)

func Test_Calc(t *testing.T) {

	var N int
	for N = 2; N <= 15; N++ {
		var i int
		var count int64
		for i = 1; i < N; i++ {
			count += int64(i) * q7.F(N, i)
		}
		t.Logf("N(%v): %v", N, count)
	}
}
