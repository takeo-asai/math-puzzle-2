package q2_test

import (
	"testing"

	"puzzle/quizzes/q2"
)

func Test_abc(t *testing.T) {
	c := q2.Main()
	t.Logf("a: %v", c)
}
