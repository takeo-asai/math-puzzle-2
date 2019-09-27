package q12_test

import (
	"testing"

	"puzzle/quizzes/q12"
)

func Test_pi(t *testing.T) {
	var n int64
	for n = 1; n < 10; n++ {
		t.Logf("n=%v, %v", n, q12.Pi(n).Value())
	}
}
